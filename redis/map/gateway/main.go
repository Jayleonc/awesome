package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Key struct {
	AppId    string                  `json:"appId"`
	KeyId    string                  `json:"keyId"`
	LesseeId string                  `json:"lesseeId"`
	Storages map[string]StoragesInfo `json:"storages"`
}

type StoragesInfo struct {
	KeyIndex  string `json:"keyIndex"`
	NodeState bool   `json:"nodeState"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.140:6379",
		Password: "smsecure123",
		DB:       0,
	})

	var key Key
	get, _ := client.HGet(context.Background(), "storage_key", "CESHI").Result()
	json.Unmarshal([]byte(get), &key)

	fmt.Println(key)
	var nodeIds []string
	var storages = key.Storages
	if len(storages) > 0 {
		for k := range storages {
			// 将每个Key(nodeId)存储起来
			nodeIds = append(nodeIds, k)
		}
	}
	fmt.Println(nodeIds)
}
