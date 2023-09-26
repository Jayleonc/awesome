package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy) //外层切片
	for x := range a {
		b := make([]uint8, dx) //里层切片
		for y := range b {
			b[y] = uint8(x*y - 1) //给里层切片里的每一个元素赋值。其中x*y可以替换成别的函数
		}
		a[x] = b //给外层切片里的每一个元素赋值
	}
	return a
}

func main() {
	s := strings.Map(func(r rune) rune {
		return r + 1
	}, "abc000")
	fmt.Println(s)
}

func show() {
	defer fmt.Println("Defer1")
	defer fmt.Println("Defer2")
	fmt.Println("Hello")
}

func addSelf(i *int) {
	*i++
}
