package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
)

func main() {
	//data := []byte("1234567812345678")

	//cipher, err := sm4.NewCipher(data)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	key, err := hex.DecodeString("hel12vde8fn128vn2957591c81rf818")
	if err != nil {
		return
	}

	context := []byte("hello world")
	ecb, err := sm4.Sm4Ecb(key, context, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ByteToString(ecb))

	sm4Ecb, err1 := sm4.Sm4Ecb(key, ecb, false)
	if err1 != nil {
		return
	}
	fmt.Println(ByteToString(sm4Ecb))
}
func ByteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	return ret
}
