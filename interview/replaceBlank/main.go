package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "This is a book"
	fmt.Println(Replace(s))

	var a Show
	a.Param = make(map[string]interface{})
	a.Param["RMB"] = 1000

	name := student{
		Name: "jay",
	}
	z(name)
}

func Replace(source string) string {

	if len(source) > 1000 {
		return source
	}

	for _, v := range source {
		if v != ' ' && !unicode.IsLetter(v) {
			return source
		}
	}

	return strings.Replace(source, " ", "%20", -1)
}

type Param map[string]interface{}

type Show struct {
	Param
}

type student struct {
	Name string
}

func z(v interface{}) {
	switch msg := v.(type) {
	case *student:
		fmt.Println(msg.Name)
	case student:
		fmt.Println(msg.Name)
	}
}
