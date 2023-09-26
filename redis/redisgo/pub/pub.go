package main

import (
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

	// 发布消息到频道
	_, err := conn.Do("PUBLISH", "channel1", "Hello, World!")
	if err != nil {
		panic(err)
	}
}
