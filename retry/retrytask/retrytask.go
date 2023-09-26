package retrytask

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type RetryFunc interface {
	Execute() error
}

type RetryTask struct {
	Id       string
	Function RetryFunc
}

type RetryCounter struct {
	count  int
	rounds int
}

var TaskQueue = make(chan RetryTask, 50)
var mu sync.Mutex
var sem chan struct{}

func RetryScheduler() {
	ticker := time.Tick(time.Duration(5) * time.Second)

	taskCount := make(map[string]*RetryCounter)
	sem = make(chan struct{}, 30)

	if true {
		for range ticker {
			// 执行任务队列中的任务
			fmt.Println("taskQueue:", len(TaskQueue))
			for len(TaskQueue) > 0 {
				log.Printf("当前任务队列长度：%v\n", len(TaskQueue))
				sem <- struct{}{}

				task := <-TaskQueue
				go func(task RetryTask) {
					defer func() {
						<-sem
					}()
					// 执行任务并处理重试逻辑
					log.Printf("任务队列重新执行定时任务：%v\n", task.Id)
					err := task.Function.Execute()

					mu.Lock()
					counter := getOrCreateCounter(taskCount, task.Id)
					if err != nil {
						handleRetry(task, counter, taskCount, err)
					} else {
						handleSuccess(task, counter)
					}
					mu.Unlock()
				}(task)
			}
		}
	}
}

// handleRetry 处理任务重试逻辑
func handleRetry(task RetryTask, counter *RetryCounter, taskCount map[string]*RetryCounter, err error) {
	counter.IncreaseCount()

	if counter.ShouldRetry() {
		log.Printf("重试任务（%s）：%v，尝试次数：%d\n", task.Id, err, counter.GetCount())
		// 启动新协程等待放回队列
		go func(task RetryTask) {
			TaskQueue <- task
			log.Printf("任务（%s）成功放回队列\n", task.Id)
		}(task)

	} else {
		log.Printf("%v 次重试后任务失败（%s）：%v\n", counter.GetCount(), task.Id, err)
		counter.ResetCount()
		counter.IncreaseRounds()

		if counter.ShouldAbandon() {
			log.Printf("放弃任务（%v），超过 3 轮\n", task.Id)
			delete(taskCount, task.Id) // 删除任务的计数器
		} else {
			go func(task RetryTask) {
				TaskQueue <- task
				log.Printf("任务（%s）成功放回队列\n", task.Id)
			}(task)
		}
	}
}

// handleSuccess 处理任务成功逻辑
func handleSuccess(task RetryTask, counter *RetryCounter) {
	log.Printf("任务队列重新执行任务成功：%v\n", task.Id)
	counter.ResetCount() // 任务成功执行，重试计数重置为 0
	counter.ResetRounds()
}

// getOrCreateCounter 获取或创建任务计数器
func getOrCreateCounter(taskCount map[string]*RetryCounter, taskID string) *RetryCounter {
	counter, exists := taskCount[taskID]
	if !exists {
		counter = &RetryCounter{}
		taskCount[taskID] = counter
	}
	return counter
}

// CreateRetryTask 创建 RetryTask 并为任务生成唯一ID
func CreateRetryTask(function RetryFunc) RetryTask {
	return RetryTask{
		Id:       generateUniqueID(),
		Function: function,
	}
}

// 生成唯一ID
func generateUniqueID() string {
	return uuid.New().String()
}

// IncreaseCount 增加计数器的尝试次数
func (c *RetryCounter) IncreaseCount() {
	c.count++
}

// GetCount 获取计数器的尝试次数
func (c *RetryCounter) GetCount() int {
	return c.count
}

// ResetCount 重置计数器的尝试次数为0
func (c *RetryCounter) ResetCount() {
	c.count = 0
}

// IncreaseRounds 增加计数器的轮数
func (c *RetryCounter) IncreaseRounds() {
	c.rounds++
}

// ResetRounds 重置计数器的轮数为0
func (c *RetryCounter) ResetRounds() {
	c.rounds = 0
}

// ShouldRetry 检查是否应该进行重试
func (c *RetryCounter) ShouldRetry() bool {
	return c.count < 5
}

// ShouldAbandon 检查是否应该放弃任务
func (c *RetryCounter) ShouldAbandon() bool {
	return c.rounds >= 3
}
