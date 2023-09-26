package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var master *redis.ClusterClient
var slave *redis.ClusterClient

func init() {

	// master/write
	mastersAddrs := []string{
		"175.178.58.198:6379",
		"175.178.58.198:6380",
		"175.178.58.198:6381",
	}

	// slave/read
	slaveAddrs := []string{
		"175.178.58.198:6382",
		"175.178.58.198:6383",
		"175.178.58.198:6384",
	}

	// 创建 Redis 集群客户端
	master = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    mastersAddrs,
		Password: "jayleonc",
	})

	slave = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    slaveAddrs,
		Password: "jayleonc",
	})
}

func main() {
	go func() {
		err := slave.Set(context.Background(), "mykey", "myvalue", 0).Err()
		if err != nil {
			fmt.Println("Error setting key:", err)
			return
		}
	}()

	time.Sleep(2)

	val, err := slave.Get(context.Background(), "mykey").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	fmt.Println("mykey:", val)
}
