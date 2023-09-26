package main

import "fmt"

var Name string
var Email string

func main() {
	fmt.Print("What's your name: ")
	fmt.Scan(&Name)
	fmt.Print("What's your email: ")
	fmt.Scan(&Email)
	fmt.Printf("Hello %v, I sent it to your email %v\n", Name, Email)
}

type I1 interface {
	M1()
}

type I2 interface {
	M8(string)
	M2()
}

type I3 interface {
	I1
	I2
	M3()
}
