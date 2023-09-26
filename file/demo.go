package main

import (
	"fmt"
	"time"
)

type student struct {
	Name string
}

func main() {
	s := student{
		Name: "Jay chou",
	}
	zhoujielun(&s)

	time.Now()
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student:
		name := msg.Name
		fmt.Println(name)
	}
}
