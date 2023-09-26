package main

import "fmt"

type Animal interface {
	SetName(string)
	GetName() string
}

type Cat struct {
	Name string
}

func (c *Cat) SetName(s string) {
	c.Name = s
	fmt.Printf("cat name is %v\n", c.Name)
}

func (c Cat) GetName() string {
	return c.Name
}

type Dog struct {
	Name string
}

func main() {
	var cat = Cat{
		Name: "mcmc",
	}

	var A Animal = &cat
	var name = A.GetName()
	fmt.Println(name)
	A.SetName("Jayleonc")
	A.SetName("Chen")

}
