package main

import (
	"fmt"
	"time"
)

// 消费-- 写入数据库操作
func consumer(cname string, ch chan string) {
	for i := range ch {
		fmt.Println("客户--", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

// 生产-- 产生数据库消息
func producer(pname string, ch chan string) {
	for i := 0; i < 4; i++ {
		s := fmt.Sprintln("面包", pname, ":", i)
		ch <- s
	}
}

func main() {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	data := make(chan string)
	go producer("生产", data)
	//go consumer("消费", data)

	time.Sleep(10 * time.Second)
	close(data)
	time.Sleep(10 * time.Second)
}
