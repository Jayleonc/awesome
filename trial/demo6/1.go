package main

import "fmt"

type User struct {
	likes int
}

func main() {

	users := make([]User, 1)
	ptrUser0 := &users[0]
	ptrUser0.likes++
	a := ptrUser0.likes
	fmt.Printf("A: %d Addr: %v\n", a, &a)
	users = append(users, User{})
	ptrUser0.likes++
	a = ptrUser0.likes
	fmt.Printf("A: %d Addr: %v\n", a, &a)

	for i := range users {
		fmt.Printf("User: %d Likes: %d Addr: %v\n", i, users[i].likes, &users[i].likes)
	}

}
