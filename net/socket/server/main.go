package main

import (
	"bufio"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:20000") // 三次握手，欢迎套接字
	if err != nil {
		return
	}
	for {
		accept, err := listen.Accept() // accept 是一个新的套接字，连接套接字
		if err != nil {
			continue
		}
		go process(accept)
	}
}

func process(accept net.Conn) {
	for {
		reader := bufio.NewReader(accept)
		var buf = make([]byte, 1024)
		read, err := reader.Read(buf)
		if err != nil {
			break
		}
		revStr := string(buf[:read])
		upper := strings.ToUpper(revStr)
		_, err = accept.Write([]byte(upper))
		if err != nil {
			break
		}
	}
}

// 快速排序 O(lonN)
// string 转 int
