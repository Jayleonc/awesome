package main

import "fmt"

func main() {
	var m = make(map[int]interface{})
	m[0] = 10
	m[1] = "go"
	m[2] = true
	fmt.Println(m)

	var a = m[2]
	_, ok := a.(int)
	if ok {
		fmt.Printf("断言成功，a的值是：%v 类型是：%T", a, a)
	} else {
		fmt.Printf("断言失败，a的值是：%v 类型是：%T", a, a)
	}

}
