package main

import "fmt"

type Callback func(x, y int) int

// 提供一个接口，让外部去实现
func test(x, y int, callback Callback) int {
	return callback(x, y)
}

// 回调函数的具体实现
func add(x, y int) int {
	return x + y
}

func main() {
	x, y := 1, 2
	fmt.Println(test(x, y, add))
}
