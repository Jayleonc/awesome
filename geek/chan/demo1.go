package main

import "fmt"

func main() {
	ch1 := make(chan interface{}, 3)
	var a interface{}
	a = 12.123
	fmt.Println(&a) // 0x140000200a8
	ch1 <- a
	fmt.Println(len(ch1))
	switch a.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	}
	close(ch1)
	for item := range ch1 {
		fmt.Println(&item) // 0x140000200c0 进入通道的是副本，而不是原本的元素
	}

}
