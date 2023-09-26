package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

//func main() {
//	bill := helloworld{"Show", "bill@email.com"}
//	bill.notify()
//
//	lisa := &helloworld{"Lisa", "lisa@email.com"}
//	lisa.notify()
//
//	bill.changeEmail("bill@newdomain.com")
//	bill.notify()
//}
