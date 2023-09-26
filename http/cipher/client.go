package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.53:1111")
	if err != nil {
		fmt.Println("coon Dial failed :", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn close failed")
		}
	}(conn)

	fmt.Println(conn.RemoteAddr().Network())
	fmt.Println(conn.RemoteAddr().String())
}
