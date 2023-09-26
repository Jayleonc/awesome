package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

type Computer struct {
}

type Camera struct {
}

func (c Computer) work(usb Usb) {
	usb.start()
	usb.stop()
}

func (p Phone) start() {
	fmt.Printf("%v 开机\n", p.Name)
}

func (p Phone) stop() {
	fmt.Printf("%v 关机\n", p.Name)
}

func (c Camera) start() {
	fmt.Println("相机开机")
}

func (c Camera) stop() {
	fmt.Println("相机关机")
}

func main() {
	var computer = Computer{}
	var phone = Phone{Name: "iPhone"}
	var camera = Camera{}
	computer.work(phone)
	computer.work(camera)
}
