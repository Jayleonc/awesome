package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("123456"))
	result := hex.EncodeToString(h.Sum(nil))
	fmt.Println(result)
}
