package main

import "fmt"

var block = "package"

func main() {
	block := "func"

	{
		block := "inner"
		fmt.Printf("%s\n", block)
	}
	fmt.Printf("%s\n", block)
	show()
}

func show() {
	fmt.Println(block)
}
