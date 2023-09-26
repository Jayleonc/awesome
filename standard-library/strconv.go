package main

import (
	"fmt"
	"strconv"
)

func main() {
	//name := "Jay"
	arr := []byte{1, 2, 3}
	appendInt := strconv.AppendInt(arr, 2, 1)
	fmt.Println(appendInt)
}
