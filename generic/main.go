package main

import "fmt"

func Print[T any](v T) {
	fmt.Printf("%T:%v\n", v, v)
}

type MyInt int

type Goer interface {
	GO()
}

type Go struct {
	Name string
}

func (g Go) GO() {
	fmt.Printf("GO %v", g.Name)
}

func main() {
	Print("ok")
	Print(MyInt(2))

	var g Goer = Go{"Ok"}

	g.GO()
	if _, ok := g.(Go); ok {
		fmt.Println("jlasdf")
	}

	switch g.(type) {
	case Go:
		fmt.Println("hahahahahah")
	default:
		fmt.Println("eeeeeeeeeee")
	}

}
