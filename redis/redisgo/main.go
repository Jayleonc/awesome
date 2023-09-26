package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:   10,  // 最大空闲连接数
		MaxActive: 100, // 最大活跃连接数
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword("jayleonc"))
		},
	}
	defer pool.Close()

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "name", "jayleonc")
	if err != nil {
		panic(err)
	}

	conn = pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GET", "name")
	if err != nil {
		panic(err)
	}

	fmt.Println("value:", string(reply.([]byte)))

	go func() {
		// 订阅频道
		conn.Do("SUBSCRIBE", "channel1", "channel2")

		// 接收消息
		receive, err := conn.Receive()
		if err != nil {
			panic(err)
		}

		switch rep := receive.(type) {
		case redis.Message:
			fmt.Println("Received message:", string(rep.Data))
		case redis.Subscription:
			fmt.Println("Subscription:", rep.Kind, rep.Channel, rep.Count)
		}
	}()
}
