package main

import (
	"fmt"
	_ "github.com/imroc/biu"
)

func main() {

	//var b = []byte("12 8 21 68 -39 62 -76 86")
	//// 将 []byte 转成 16 进制字符串
	//toString := hex.EncodeToString(b)
	//fmt.Println(toString)
	//
	//fmt.Println("--------------------------------")
	//str := "01000040000000000000007700000000000000000100021015b4492130215f5800000000c000188000770000000c00144256cb0a98b0195d61f1422e49b14d97fa7d9937"
	//
	//// 将 16 进制字符串转成 []byte
	//decodeString, err := hex.DecodeString(str)
	//if err != nil {
	//	return
	//}
	//var bytes = decodeString
	//fmt.Println("---------------------")
	//fmt.Printf("%T\n", bytes)
	//fmt.Println(bytes)
	//
	//binaryString := biu.BytesToBinaryString(b)
	//fmt.Println(binaryString)

	var arr = []byte("1234567")
	for i := range arr {
		fmt.Println(arr[i])
	}
	htonl := Htonl(&arr, 0, 10)
	fmt.Println(htonl)

}

func Htonl(data *[]byte, offset int, value int) *[]byte {
	(*data)[offset] = byte(value >> 24 & 255)
	(*data)[offset+1] = byte(value >> 16 & 255)
	(*data)[offset+2] = byte(value >> 8 & 255)
	(*data)[offset+3] = byte(value >> 0 & 255)
	return data
}
