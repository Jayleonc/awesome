package main

import "fmt"

func main() {
	RecoverTestFail(5)
}

func RecoverTestFail(num int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover: %s\n", err)
		}
	}()
	if num%5 == 0 {
		panic("本协程请求出错")
	}
	go AnotherRoutinePrint(num)
}

func AnotherRoutinePrint(num int) {
	if num%4 == 0 {
		panic("其他协程请求又出错了")
	}
	fmt.Printf("%d\n", num)
}
