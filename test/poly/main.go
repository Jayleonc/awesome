package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Change interface {
	Update()
	Create()
	Delete()
}

type Endpoint struct {
	ID   string
	Ip   string
	Port string
}

var EndpointItem *Endpoint

func (e Endpoint) Create() {
	message := map[string]interface{}{
		"Endpoint": e,
	}
	Pub(message)
}
func (e Endpoint) Update() {}
func (e Endpoint) Delete() {}
func (e Endpoint) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}
func AdminChangeChangeEndpoint() {
	i := Endpoint{
		ID:   "123456",
		Ip:   "192.168.1.140",
		Port: "8989",
	}
	EndpointItem = &i
}

func main() {
	AdminChangeChangeEndpoint()
	EndpointItem.Create()
}

// Pub 发布消息
func Pub(message map[string]interface{}) {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "jayleonc",
		DB:       0,
	})

	// 发布消息
	streamID, err := client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "jayStream", // 要发布到的 Stream 的名称
		Values: message,     // 消息内容
	}).Result()
	if err != nil {
		fmt.Println("发布消息出错:", err)
		return
	}
	fmt.Println("消息已发布，StreamID:", streamID)
}
