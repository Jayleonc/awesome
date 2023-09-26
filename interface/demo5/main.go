package main

import "fmt"

type admin struct {
	id string
	user
}

type user struct {
	id    int    // 8
	isMen bool   // 1
	email string // 16
}

func main() {
	arr := []int{1, 2, 3}
	var newArr []*int
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}

	var arr1 = make([]int, len(arr))
	copy(arr1, arr)
	fmt.Println(arr1)
}
