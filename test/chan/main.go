package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func writeLog(i int, ch chan int) {
	file, err := os.OpenFile("src/github.com/jayleonc/test/chan/config.txt", os.O_RDWR|os.O_APPEND, 0776)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("写入文件--" + strconv.Itoa(i) + "\n")
	ch <- i
}

func main() {
	fmt.Println("开始------")
	start := time.Now().UnixNano()

	ch := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		go writeLog(i, ch)
	}
	for j := 0; j < 5; j++ {
		fmt.Println(<-ch)
	}
	end := time.Now().UnixNano()
	fmt.Println("结束------", (end-start)/1e6)
}
