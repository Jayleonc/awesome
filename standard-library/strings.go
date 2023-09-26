package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = strings.Builder{}
	str.WriteByte('a')
	str.WriteString("->hahahahah")
	fmt.Println(str.String())

	a := "a"
	b := "c"
	if strings.Compare(a, b) != 0 {
		fmt.Println("不相等")
	}

	clone := strings.Clone(a)
	fmt.Println(clone)

	if strings.Contains(str.String(), "ahahahah") == true {
		fmt.Println("包含")
	}
}
