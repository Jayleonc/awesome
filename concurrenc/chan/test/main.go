package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch = make(chan interface{}, 10)
	//go client(ch)
	//quit := make(chan bool)
	//go server(ch, quit)
	//<-quit
	//fmt.Println("quit")
	data := make(chan int)
	shutdown := make(chan int)
	close(shutdown)
	close(data)

	select {
	case s := <-shutdown:
		fmt.Println(s)
		fmt.Println("CLOSED, ")
	case data <- 1:
		fmt.Print("HAS WRITTEN, ")
	default:
		fmt.Print("DEFAULT, ")
	}
}

func server(ch chan any, q chan bool) {
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			}
		case <-time.After(1000):
			q <- true
		}
	}
}

func client(ch chan any) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
