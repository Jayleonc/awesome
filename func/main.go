package main

import "fmt"

// 这创建了类型定义，创建了一个自定义类型 add，表示为 func(int, int) int 的函数类型
type add func(int, int) int

// 这是一个类型别名，不会创建新的类型，而是为该函数类型创建一个可替代的名称 add2
type add2 = func(int, int) int

func addNumber(a, b int) int {
	return a + b
}

func addOneAndTwo(a, b int) int {
	c := a + b
	if c != 3 {
		fmt.Println("出错啦！！")
		return 0
	}
	return a + b
}

func main() {
	a := add(addNumber)
	i := a(1, 2)
	fmt.Println(i)

}
