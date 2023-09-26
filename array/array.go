package main

import "fmt"

func main() {
	myArr1 := [10]int{1, 2, 3, 4}
	myArr2 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myArr1); i++ {
		fmt.Println(myArr1[i])
	}
	fmt.Println("------------")
	printArr(myArr2)
}

func printArr(myArr [4]int) {
	for i, i2 := range myArr {
		fmt.Println("index = ", i, ", value = ", i2)
	}
}
