package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
)

func main() {
	data := []byte("hello world")
	hw := sm3.New()
	hw.Write(data)
	hash := hw.Sum(nil)
	toString := ByteToString(hash)

	hw2 := sm3.New()
	data2 := []byte("hello world1")
	hw2.Write(data2)
	hash2 := hw2.Sum(nil)
	toString2 := ByteToString(hash2)
	if toString == toString2 {
		fmt.Println("相等")
	}

}

func ByteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	fmt.Println("ret = ", ret)
	return ret
}
