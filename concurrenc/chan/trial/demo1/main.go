package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	hello := "Hello, world"
	helloByte := []byte(hello)
	var wg sync.WaitGroup
	wg.Add(2)
	str := make(chan byte, len(helloByte))
	for i := range helloByte {
		str <- helloByte[i]
	}
	close(str)

	ch1 := make(chan byte, len(str))
	ch2 := make(chan byte, len(str))

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		for {
			select {
			case v, ok := <-ch1:
				if ok {
					fmt.Printf("ch1 : %v\n", string(v))
				} else {
					close(ch1)
				}
			case v, ok := <-ch2:
				if ok {
					fmt.Printf("ch2 : %v\n", string(v))
				} else {
					close(ch2)
				}
			default:
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for b := range str {
			select {
			case ch1 <- b:
			case ch2 <- b:
			}
		}
	}()

	wg.Wait()
}
