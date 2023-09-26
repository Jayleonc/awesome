package main

import (
	"unsafe"
)

const size = 100

func main() {
	a := make([]int, 1)    // 无法被一个执行栈装下，即便没有返回，也会直接在堆上分配；
	b := make([]int, 8192) // 对象能够被一个执行栈装下，变量没有返回到栈外，进而没有发生逃逸。
	// 一个栈占有 8291 = 24 个字节 = 192 位
	println(a)
	println(unsafe.Sizeof(b)) // 返回字节大小 24 字节

	//strings := make([]int, 3)
	//strings[0] = 1
	//strings[0] = 2
	//strings[0] = 3
	//println(&strings[0])
	//println(&strings[1])
	//println(&strings[2])
	//
	//println("---------------")
	//for i := range strings {
	//	println(i, strings[i])
	//}
	//
	//println("---------------")
	//for i, fruit := range strings {
	//	fmt.Println(i, fruit)
	//}
}
