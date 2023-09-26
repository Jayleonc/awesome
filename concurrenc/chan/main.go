package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	wg      sync.WaitGroup
	channel = make(chan int, 10000)
)

func main() {
	//先写满一个channel
	for i := 0; i < 5; i++ {
		channel <- i
	}
	wg.Add(2)
	go HandleSignals()
	go func() {
		defer wg.Done()
		for {
			select {
			case num, ok := <-channel:
				if !ok {
					fmt.Println("!ok")
					return
				}
				fmt.Println("======", num)
				//每次从channel取值后sleep 1秒，方便我们分析
				time.Sleep(time.Duration(num) * time.Second)
			}
		}
	}()
	wg.Wait()
}
func HandleSignals() {
	defer wg.Done()

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("Exiting, please wait...")
			close(channel)
			return
		}
	}
}
