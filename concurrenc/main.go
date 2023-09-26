package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	//wg 用来等待程序完成
	//add 2 表示要等待两个 goroutine
	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")
	go printPrime("C")

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")

	// 声明一个匿名函数，并创建一个 goroutine
	//	go func() {
	//		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	//		defer wg.Done()
	//
	//		for count := 0; count < 3; count++ {
	//			for char := 'a'; char < 'a'+26; char++ {
	//				fmt.Printf("%c ", char)
	//			}
	//		}
	//	}()

	// 声明一个匿名函数，并创建一个 goroutine
	//	go func() {
	//		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	//		defer wg.Done()
	//
	//		for count := 0; count < 3; count++ {
	//			for char := 'A'; char < 'A'+26; char++ {
	//				fmt.Printf("%c ", char)
	//			}
	//		}
	//	}()

	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// 显示 5000 以内的质数值
func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed ", prefix)

}
