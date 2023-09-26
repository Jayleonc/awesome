package main

import "fmt"

func main() {
	arr := []int{97, 12, 54}
	fmt.Println(arr)
	bSort(arr)
	fmt.Println(arr)
}

func bSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func swap(arr []int, i int, min int) {
	temp := arr[i]
	arr[i] = arr[min]
	arr[min] = temp
}
