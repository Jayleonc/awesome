package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		//recvStr := string(buf[:n])
		toString := hex.EncodeToString(buf[:n])
		fmt.Println("收到 client 端发来的数据：", toString)

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("err: ", err)
			return
		} // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
