package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// listenPubSubChannels listens for messages on Redis pubsub channels. The
// onStart function is called after the channels are subscribed. The onMessage
// function is called for each message.
func listenPubSubChannels(ctx context.Context, redisServerAddr string,
	onStart func() error,
	onMessage func(channel string, data []byte) error,
	channels ...string) error {
	// A ping is set to the server with this period to test for the health of
	// the connection and server.
	const healthCheckPeriod = time.Minute
	password := redis.DialPassword("jayleonc")
	c, err := redis.Dial("tcp", redisServerAddr,
		// Read timeout on server should be greater than ping period.
		redis.DialReadTimeout(healthCheckPeriod+10*time.Second),
		redis.DialWriteTimeout(10*time.Second), password)
	if err != nil {
		return err
	}
	defer c.Close()

	psc := redis.PubSubConn{Conn: c}

	if err := psc.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
		return err
	}

	done := make(chan error, 1)

	// Start a goroutine to receive notifications from the server.
	go func() {
		for {
			switch n := psc.Receive().(type) {
			case error:
				done <- n
				return
			case redis.Message:
				if err := onMessage(n.Channel, n.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				switch n.Count {
				case len(channels):
					// Notify application when all channels are subscribed.
					if err := onStart(); err != nil {
						done <- err
						return
					}
				case 0:
					// Return from the goroutine when all channels are unsubscribed.
					done <- nil
					return
				}
			}
		}
	}()

	ticker := time.NewTicker(healthCheckPeriod)
	defer ticker.Stop()
loop:
	for {
		select {
		case <-ticker.C:
			// Send ping to test health of connection and server. If
			// corresponding pong is not received, then receive on the
			// connection will timeout and the receive goroutine will exit.
			if err = psc.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			break loop
		case err := <-done:
			// Return error from the receive goroutine.
			return err
		}
	}

	// Signal the receiving goroutine to exit by unsubscribing from all channels.
	if err := psc.Unsubscribe(); err != nil {
		return err
	}

	// Wait for goroutine to complete.
	return <-done
}

func publish() {
	c, err := dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	if _, err = c.Do("PUBLISH", "c1", "hello"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err = c.Do("PUBLISH", "c2", "world"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err = c.Do("PUBLISH", "c1", "goodbye"); err != nil {
		fmt.Println(err)
		return
	}
}

// This example shows how receive pubsub notifications with cancelation and
// health checks.
func ExamplePubSubConn() {
	redisServerAddr, err := serverAddr()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	err = listenPubSubChannels(ctx,
		redisServerAddr,
		func() error { go publish(); return nil },
		func(channel string, message []byte) error {
			fmt.Printf("channel: %s, message: %s\n", channel, message)
			if string(message) == "goodbye" {
				cancel()
			}
			return nil
		},
		"c1", "c2")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Output:
	// channel: c1, message: hello
	// channel: c2, message: world
	// channel: c1, message: goodbye
}

func dial() (redis.Conn, error) {
	password := redis.DialPassword("jayleonc")
	conn, t := redis.Dial("tcp", "localhost:6379", password)
	if t != nil {
		return nil, t
	}
	return conn, nil
}

// serverAddr wraps DefaultServerAddr() with a more suitable function name for examples.
func serverAddr() (string, error) {
	return "localhost:6379", nil
}

func main() {
	ExamplePubSubConn()
}
