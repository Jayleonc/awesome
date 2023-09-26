package main

import (
	"fmt"
	"os"
	"time"
)

var query = "test"
var matches int

func main() {
	start := time.Now()
	search("/Users/jayleonc/")
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func search(path string) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				matches++
			}
			if file.IsDir() {
				search(path + name + "/")
			}
		}
	}
}
