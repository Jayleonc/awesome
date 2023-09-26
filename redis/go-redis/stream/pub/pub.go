package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

const ServicesKey = "Service"

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "jayleonc",
		DB:       0,
	})

	// 发布消息
	for i := 0; i < 1; i++ {
		// 发布消息到 Stream
		streamID, err := client.XAdd(&redis.XAddArgs{
			Stream: "jayStream", // 要发布到的 Stream 的名称
			Values: map[string]interface{}{
				"key":     "Service",                                // key
				"content": "Hello, Redis Stream " + strconv.Itoa(i), // 消息内容

			},
		}).Result()
		if err != nil {
			fmt.Println("发布消息出错:", err)
			return
		}
		fmt.Println("消息已发布，StreamID:", streamID)
	}
	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}
