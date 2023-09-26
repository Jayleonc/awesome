package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func main() {
	// 对sha256算法进行hash加密,key随便设置
	hash := hmac.New(sha1.New, []byte("abc123")) // 创建对应的sha256哈希加密算法
	msg := "这里是要加密的数据"
	hash.Write([]byte(msg))
	b := hash.Sum(nil)

	hash2 := hmac.New(sha1.New, []byte("abc123")) // 创建对应的sha256哈希加密算法
	msg2 := "这里是要加密的数据"
	hash2.Write([]byte(msg2))
	b2 := hash2.Sum(nil)
	equal := hmac.Equal(b, b2)
	fmt.Println(equal)

}
