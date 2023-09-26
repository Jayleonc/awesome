package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	var end *time.Time
	sec := int(n.Sub(*end).Seconds())
	if end != nil && sec < 4 {
		fmt.Println("123")
	}
}
