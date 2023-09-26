package main

import (
	"fmt"
	"unsafe"
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
	mutexWaiterShift1
)

type Person struct {
	Name    string
	Age     int
	Address string
	Phone   map[string]string
}

func main() {
	var x = "Hello"
	fmt.Println(&x)
	fix(x)

	var i int64 = 1
	fmt.Println(unsafe.Sizeof(i))

	var sl1 []int
	var sl2 = []int{1, 2, 3}
	sl1 = append(sl1, sl2...)
	fmt.Println(&sl1[1])

	mymap := new(int)
	*mymap = 1
	fmt.Println(*mymap)

	var p Person // 创建一个 Person 类型的零值
	p.Phone = make(map[string]string)
	p.Phone["1"] = "182" // 此 p.Phone 是 nil 的映射，没有分配内存

	type O struct {
		l int64 // 8
		n int64 // 8
	}
	var o O
	fmt.Printf("Size: %v", unsafe.Sizeof(o))

}

func fix(x string) {
	fmt.Println(&x)
}
