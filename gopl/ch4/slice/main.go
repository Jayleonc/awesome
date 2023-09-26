package main

import (
	"fmt"
)

func main() {
	months := [...]string{
		1:  "一月",
		2:  "二月",
		3:  "三月",
		4:  "四月",
		5:  "五月",
		6:  "六月",
		7:  "七月",
		8:  "八月",
		9:  "九月",
		10: "十月",
		11: "十一月",
		12: "十二月",
	}
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", months, months, len(months), cap(months))

	Q1 := months[1:4]
	Q2 := months[4:7]
	Q3 := months[7:10]
	Q4 := months[10:]
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", Q1, Q1, len(Q1), cap(Q1))
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", Q2, Q2, len(Q2), cap(Q2))
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", Q3, Q3, len(Q3), cap(Q3))
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", Q4, Q4, len(Q4), cap(Q4))

	arr := make([]int, 8, 12)
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", arr, arr, len(arr), cap(arr))
	arr = append(arr, 1, 2, 3, 4)
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", arr, arr, len(arr), cap(arr))

	arr2 := make([]int, 0, 12)
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", arr2, arr2, len(arr2), cap(arr2))

	if len(arr2) == 0 {
		fmt.Println("arr2 values is nil")
	}

	i := []int{1, 2, 3, 4}
	x := []int{5, 6, 7, 8}
	i = append(i, x...)
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", i, i, len(i), cap(i))

	// pop，弹出最后一个元素
	i = i[:len(i)-1]
	fmt.Printf("type: %T value: %v len: %d cap: %d\n", i, i, len(i), cap(i))

	// 栈顶
	s := i[len(i)-1]
	fmt.Println(s)
}
