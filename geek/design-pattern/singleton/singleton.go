package singleton

import "fmt"

type Singleton struct {
	Name string
	age  int
}

var ins *Singleton

func GetInsOr() *Singleton {
	ins = &Singleton{}
	return ins
}

func main() {
	or := GetInsOr()
	fmt.Println(or)
}
