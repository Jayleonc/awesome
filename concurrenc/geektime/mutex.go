package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func main() {
	// Mutex 就提供了两个方法 Lock 和 Unlock：进入临界区之前调用 Lock，退出临界区的时候调用 UnLock

	// 互斥锁保护计数器
	//var mu sync.Mutex
	var counter Counter
	// 计数器的值
	//var count = 0
	// 使用 WaitGroup ，用来确保所有的 goroutine 都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动 10 个 goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())

}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
