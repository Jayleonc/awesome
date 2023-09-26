package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := "Hello,world"
	helloStr := []byte(hello)
	hechan := make(chan byte, 100)
	for i := range helloStr {
		hechan <- helloStr[i]
	}
	close(hechan)

	count := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)

	threeChan := make(chan int)

	go func() {
		defer func() {
			wg.Done()
			if !isChanClose(threeChan) {
				threeChan <- 10
			}
		}()
		for {
			if i, ok := <-count; ok {
				if v, sok := <-hechan; sok {
					fmt.Printf("chan1: %c\n", v)
				} else {
					close(count)
					break
				}
				count <- i
			} else {
				break
			}
		}
	}()

	go func() {
		defer func() {
			wg.Done()
			if !isChanClose(threeChan) {
				threeChan <- 10
			}
		}()
		for {
			if i, ok := <-count; ok {
				if v, sok := <-hechan; sok {
					fmt.Printf("chan2: %c\n", v)
				} else {
					close(count)
					break
				}
				count <- i
			} else {
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if i, ok := <-threeChan; ok {
				fmt.Println("第三个 goroutine 执行啦！！！", i)
				close(threeChan)
			} else {
				break
			}
		}
	}()

	count <- 1
	wg.Wait()
}

func isChanClose(ch chan int) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}
