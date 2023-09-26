package main

import (
	"fmt"
	"unsafe"
)

type Usb interface {
	start()
	stop()
}

// 如果接口里有方法，必须通过结构体或自定义类型来实现方法

type Phone struct {
	Name string
}

func (p Phone) start() { // 值接收者
	fmt.Printf("%v 开机\n", p.Name)
}

func (p Phone) stop() {
	fmt.Printf("%v 关机\n", p.Name)
}

func main() {
	p := Phone{
		Name: "iPhone",
	}
	var usb Usb = p
	usb.start()
	usb.stop()

	var i item

	fmt.Println(unsafe.Sizeof(i))

}

type item struct {
	name string
}
