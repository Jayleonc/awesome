package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:   10,  // 最大空闲连接数
		MaxActive: 100, // 最大活跃连接数
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "175.178.58.198:6379", redis.DialPassword("jayleonc"))
		},
	}
	defer pool.Close()
	go func() {
		conn := pool.Get()
		defer conn.Close()
		// 发布消息到频道
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			_, err := conn.Do("PUBLISH", "channel1", "Hello, World "+strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
		}
		_, err := conn.Do("PUBLISH", "channel1", "bye")
		if err != nil {
			panic(err)
		}
	}()

	// 创建一个通道用于通知订阅结束
	done := make(chan bool)

	go func() {
		conn := pool.Get()
		defer conn.Close()
		// 订阅频道
		subConn := redis.PubSubConn{Conn: conn}
		err := subConn.Subscribe("channel1")
		if err != nil {
			panic(err)
		}

		// 接收消息
		for {
			switch rep := subConn.Receive().(type) {
			case redis.Message:
				fmt.Println("Received message:", string(rep.Data))
				if string(rep.Data) == "bye" {
					conn.Do("UNSUBSCRIBE", "channel1")
					done <- true
					return
				}
			case redis.Subscription:
				fmt.Println("Subscription:", rep.Kind, rep.Channel, rep.Count)
				if rep.Count == 0 {
					fmt.Println()
					conn.Do("UNSUBSCRIBE", "channel1")
					done <- true
					return
				}
			}
		}
	}()

	<-done
	fmt.Println("Done")
}
