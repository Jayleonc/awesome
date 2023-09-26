package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var isOk bool
	fmt.Println(isOk)

	var num rune = 'C'
	fmt.Println(unsafe.Sizeof(num))

	var floatNum float32
	fmt.Println(unsafe.Sizeof(floatNum))

	var name string
	fmt.Println(unsafe.Sizeof(name))
}
