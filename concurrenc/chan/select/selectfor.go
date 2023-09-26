package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	for {
		if len(intChan) <= 0 {
			break
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Printf("End.\n")
				break
			}
			fmt.Printf("Received: %v\n", e)
		}
	}
}
