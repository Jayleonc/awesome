package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("username is %v, email is %v\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("username is %v, email is %v\n", a.name, a.email)
}

func main() {

	bill := user{"bill", "bill@email.com"}
	sentNotification(&bill)

	lisa := admin{"lisa", "lisa@email.com"}
	sentNotification(&lisa)
}

func sentNotification(n notifier) {

	n.notify()
}
