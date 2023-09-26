package main

import "fmt"

type User struct {
	Name  string
	Phone string
	//helloworld  User // 如此定义，编译器无法知道 User 结构体的大小，因为这样写，将是无限，这是无效的递归
	user *User
}

type Admin struct {
	role string
	User
}

type Home struct {
	address string
	Admin
}

func main() {
	M := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	if m, ok := M[4]; !ok {
		fmt.Println("该元素不存在，")
	} else {
		fmt.Println(m)
	}

	var home Home
	home.Name = "Jayleonc"

}
