package main

import (
	"encoding/hex"
	"fmt"
)

func arrayTest() {
	// 求出数组里元素的和以及这些元素的平均值，分别用 for 和 for-range 实现
	var arr = [...]int{12, 3, 41, 235, 56, 2, 9}
	var sum = 0
	//for i := 0; i <= len(arr)-1; i++ {
	//	sum += arr[i]
	//}

	for _, i := range arr {
		sum += i
	}
	fmt.Printf("sum = %v 平均值 = %.2f", sum, float64(sum)/float64(len(arr)))
}

func main() {
	//var password = []byte("12345678")
	//var sessionId = []byte("12345678")
	//var msg = []byte("text")
	//hmacSpecKey := pbkdf2.key(password, sessionId, 4, 20, sha1.New)
	//encodedStr := hex.EncodeToString(hmacSpecKey)
	//fmt.Printf("%v\n", encodedStr)
	//hash := hmac.New(sha1.New, hmacSpecKey)
	//hash.Write(msg)
	//pHmac := hash.Sum(nil)
	//toString := hex.EncodeToString(pHmac)
	//fmt.Println("pHmac:", toString)
	//var b = []byte("-79, 120, 13, 45, 122, 11, 91, 83")
	var b = []byte("1, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 119, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 16, -89, -16, -108, -96, -107, -33, -21, 125, 0, 0, 0, 0, -64, 0, 24, -128, 0, 119, 0, 0, 0, 12, 0, 20, 1, -33, -25, 62, 3, 4, 86, 26, -87, -54, 32, 124, 7, 102, 68, -13, 86, 8, -13, 3, -89, -16, -108, -96, -107, -33, -21, 125")
	toString := hex.EncodeToString(b)
	//fmt.Println(toString)
	//str := "312c20302c20302c2036342c20302c20302c20302c20302c20302c20302c20302c203131392c20302c20302c20302c20302c20302c20302c20302c20302c20312c20302c20322c2031362c202d38392c202d31362c202d3130382c202d39362c202d3130372c202d33332c202d32312c203132352c20302c20302c20302c20302c202d36342c20302c2032342c202d3132382c20302c203131392c20302c20302c20302c2031322c20302c2032302c20312c202d33332c202d32352c2036322c20332c20342c2038362c2032362c202d38372c202d35342c2033322c203132342c20372c203130322c2036382c202d31332c2038362c20382c202d31332c20332c202d38392c202d31362c202d3130382c202d39362c202d3130372c202d33332c202d32312c20313235"
	decodeString, err := hex.DecodeString(toString)
	if err != nil {
		return
	}
	var sessionId = decodeString

	fmt.Println(sessionId)

}
