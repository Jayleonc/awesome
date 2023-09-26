package main

import (
	"fmt"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {
	dog := Dog{}
	dog.SetName("Chen")
	fmt.Println(dog.name)
	fmt.Println(&dog)
}
