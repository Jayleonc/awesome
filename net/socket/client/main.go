package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listen, err := net.Dial("tcp", "localhost:20000") // 发起连接，执行三次握手
	if err != nil {
		return // 三次握手失败
	}
	for {
		var b = make([]byte, 1024)

		var input string
		fmt.Scanln(&input)
		input = strings.Trim(input, "\r\n")
		b = []byte(input)
		_, err := listen.Write(b[:])
		if err != nil {
			break
		}

		var rev = make([]byte, 1024)
		read, err := listen.Read(rev[:])
		if err != nil {
			break
		}
		fmt.Println(input, "to upper:", string(rev[:read]))
	}
}
