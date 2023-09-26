package main

import "fmt"

func main() {
	DrawTriangle(5, '*')
	fmt.Println()
	DrawTriangle(10, '#')
}

func DrawTriangle(n int, c rune) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for z := 1; z <= 2*i-1; z++ {
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}
