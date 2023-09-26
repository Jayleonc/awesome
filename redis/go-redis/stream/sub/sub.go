package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var StreamName = "jayStream"
var ConsumerGroup = "jayGroup"

const ServicesKey = "Service"

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "jayleonc",
		DB:       0,
	})

	done := make(chan bool)

	// 启动消费者
	go func() {
		consumer := "consumer1" // 消费者名称
		for {
			// 从 Stream 中读取消息
			result, err := client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
				Group:    ConsumerGroup,             // 消费者组名称
				Consumer: consumer,                  // 消费者名称
				Streams:  []string{StreamName, ">"}, // Stream 名称和读取位置
				Count:    2,                         // 一次读取的消息数量
				Block:    0,                         // 阻塞时间，0 表示无限阻塞
			}).Result()

			if err == redis.Nil {
				continue // 没有消息继续循环
			} else if err != nil {
				fmt.Println("读取消息出错:", err)
				done <- true
				break
			}

			// 处理消息
			for _, stream := range result {
				for _, message := range stream.Messages {
					// 消息类型判断
					switch message.Values["key"] {
					case ServicesKey:
						fmt.Println(ServicesKey + ", 被打印了")
						i := message.Values["content"]
						fmt.Println(i)
					default:
						fmt.Println(message.ID)
						fmt.Println(message.Values)
					}

					// 确认消息消费
					_, err := client.XAck(context.Background(), StreamName, ConsumerGroup, message.ID).Result()
					if err != nil {
						log.Fatalf("Failed to acknowledge message: %s", err)
					}

					// 删除消息
					_, err = client.XDel(context.Background(), StreamName, message.ID).Result()
					if err != nil {
						log.Fatalf("Failed to delete message: %s", err)
					}
				}
			}
		}
	}()

	// 等待消费者处理消息
	<-done
	fmt.Println("\nDone")
}
