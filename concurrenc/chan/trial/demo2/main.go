package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := "hello world"
	helloByte := []byte(hello)
	var wg sync.WaitGroup
	wg.Add(2)
	str := make(chan byte, len(helloByte))
	for i := range helloByte {
		str <- helloByte[i]
	}
	close(str)

	count := make(chan int)
	go func() {
		defer wg.Done()
		for {
			c, ok := <-count
			if ok {
				s, sok := <-str
				if sok {
					fmt.Printf("G1 %c\n", s)
				} else {
					close(count)
					return
				}
			} else {
				return
			}
			count <- c
		}
	}()

	go func() {
		defer wg.Done()
		for {
			c, ok := <-count
			if ok {
				s, sok := <-str
				if sok {
					fmt.Printf("G2 %c\n", s)
				} else {
					close(count)
					return
				}
			} else {
				return
			}
			count <- c
		}
	}()

	count <- 1
	wg.Wait()
}
