package main

import "fmt"

var intChan1 = make(chan int, 5)
var intChan2 = make(chan int, 5)
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("Default!")
	}
}

func getNumber(i int) int {
	return numbers[i]
}

func getChan(i int) chan int {
	return channels[i]
}
