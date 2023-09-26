package main

import (
	"fmt"
	"github.com/Jayleonc/awesome/retry/retrytask"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {

	go retrytask.RetryScheduler()

	// 模拟添加50个任务到队列
	for i := 0; i < 10000; i++ {
		task := retrytask.CreateRetryTask(&RetryHello{})
		if rand.Float64() < 0.8 {
			retrytask.TaskQueue <- task
		}
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

type RetryHello struct {
	Name string
}

func (r *RetryHello) Execute() error {
	// 模拟任务执行耗时1秒钟
	time.Sleep(900 * time.Millisecond)

	// 模拟任务成功或失败，成功概率为50%
	if rand.Float64() <= 0.5 {
		logSuccess()
		return nil // 任务执行成功
	}

	// 任务执行失败，返回错误
	err := fmt.Errorf("task execution failed")
	logError(err)
	return err
}

func logSuccess() {
	fmt.Println("[Success]: Task executed successfully")
	fmt.Println("「TaskQueue」Size ", len(retrytask.TaskQueue))
}

func logError(err error) {
	fmt.Printf("[Error]: %v\n", err)
	fmt.Println("「TaskQueue」Size ", len(retrytask.TaskQueue))
}
