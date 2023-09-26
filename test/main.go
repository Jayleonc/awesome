package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//c := C{10, &Point{1, 2}}
	//fmt.Println(c.point.x)
	//
	//var b bytes.Buffer
	//
	//b.Write([]byte("Hello"))
	//
	//_, err := fmt.Fprintf(&b, "World!")
	//if err != nil {
	//	return
	//}
	//
	//_, err = io.Copy(os.Stdout, &b)
	//if err != nil {
	//	return
	//}
	length := 8

	nums := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		nums[i] = byte(rand.Intn(200) - 100)
	}
	fmt.Println(nums)

	round := int(math.Round(10.53) / 2.3)
	fmt.Println(round)
}
