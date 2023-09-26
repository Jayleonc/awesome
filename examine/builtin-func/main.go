package main

import "fmt"

func main() {
	i := new(int)
	fmt.Printf("%v\n", *i)
	ch := make(chan string, 5)
	ch <- "Hello"
	ch <- "HaiCoder"
	ch <- "Python"
	close(ch)
	for {
		if msg, ok := <-ch; ok == false {
			fmt.Println("chan is closed")
			break
		} else {
			fmt.Println("Msg =", msg)
		}
	}

	var arr = make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)
	arr2 := []int{4, 5, 6, 7, 8}
	copy(arr1, arr2)
	fmt.Println(arr1)
	fmt.Println(arr2)
	delete(arr, 2)
	fmt.Println(arr)
	defer func() {
		a := recover()
		fmt.Println("a = ", a)
	}()
	panic("ok")

}
