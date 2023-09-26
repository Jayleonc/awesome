package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) Lisa() {
	fmt.Printf("name = %v\n", u.name)
	fmt.Printf("email = %v\n", u.email)
}

func (u User) Show() {
	fmt.Printf("name = %v\n", u.name)
	fmt.Printf("email = %v\n", u.email)
}

func (u *User) ChangeEmail(email string) {
	u.email = email
}

func main() {
	// bill 不是包名，而是变量名
	bill := User{"bill", "bill@163.com"}
	// 使用 bill 的值作为接收者进行调用，方法 Show 接收到 bill 的值的一个副本
	bill.Show()

	// 声明了一个名为 lisa 的指针变量，并使用给定的名字和电子邮件地址做初始化
	lisa := &User{"Lisa", "lisa@qq.com"}
	// 使用这个指针变量来调用 Show 方法，在背后，go 语言所做的事情：(*lisa).Show()
	// Show 操作的只是一个副本，只不过这次操作的是从 lisa 指针指向的值的副本。
	lisa.Show()

	// 值接收者使用值的副本来调用方法，而指针接受者使用实际值来调用方法
	lisa.ChangeEmail("lisa@email.com")
	lisa.Show()

	// 使用值变量调用使用指针接收者声明的方法 ChangeEmail
	// 为了适应这种方法调用，在背后 go 语言所做的事情：(&bill).ChangeEmail
	bill.ChangeEmail("bill@email.com")
	bill.Show()
}
