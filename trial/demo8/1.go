package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())
}

//go:noinline
func e(slice []string, str string, i int) {
	panic("Want stack trace")
}
