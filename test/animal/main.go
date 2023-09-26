package main

import "fmt"

type Eater interface {
	eat()
}

type Person struct {
	name string
	age  uint8
}

func (p Person) eat() {
	//TODO implement me
	fmt.Println("My name", p.name, ", I am", p.age, " I am young")
}

type Dog struct {
	name string
	age  uint8
}

func (d Dog) eat() {
	//TODO implement me
	fmt.Println("My name", d.name, ", I am", d.age, " I am old")
}

func main() {
	//eaters := []Eater{
	//	Person{
	//		name: "Jay",
	//		age:  12,
	//	},
	//	Dog{
	//		name: "gigi",
	//		age:  12,
	//	},
	//}
	//
	//for i := range eaters {
	//
	//	eaters[i].eat()
	//}
	mapp := make([]map[int]int, 10)
	generics := reverseWithGenerics(mapp)
	fmt.Println(generics)
}
func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}
