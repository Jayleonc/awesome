package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	format := "2006-01-02 15:04:05"
	//timestamp := int64(1679121734)
	//unix := time.Unix(timestamp, 0).Format(format)
	time.Now().Format(format)
	time := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	fmt.Println(time)
}
