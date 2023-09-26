package main

import (
	"fmt"
	"strings"
)

func main() {
	//var MyMap = make(map[int]string, 5)
	//MyMap[1] = "ok"
	//fmt.Println(MyMap)

	s := strings.Map(func(r rune) rune {
		return 3
	}, "ok")
	fmt.Println(s[1])
	i := x()
	m := i()
	fmt.Println(m)

}

func x() func() int {
	return func() int {
		return 100
	}
}
