package main

import (
	"fmt"
	"time"
)

func gA(a <-chan int) {
	val := <-a
	fmt.Println("a data:", val)
	return
}

func gB(b <-chan int) {
	val := <-b
	fmt.Println("b data:", val)
	return
}

func gC(c <-chan int) {
	val := <-c
	fmt.Println("c data:", val)
	return
}

func main() {
	ch := make(chan int)
	go gC(ch)
	ch <- 2
	close(ch)
	val := <-ch
	fmt.Println("recv data from closed chan :", val)
	//go gA(ch)
	//go gB(ch)
	time.Sleep(time.Second)
}
