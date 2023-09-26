package main

import "fmt"

type myInt int

const n = 1
const z = 188888

func main() {
	var m myInt = 1
	fmt.Println(m + n)

	var k int32 = 1
	j := k + z
	fmt.Println(j)

	// 变量
	var name = "jayleonc"
	fmt.Println(name)

	// 常量
	const age = 12
	fmt.Println(age)

	// 枚举
	type PolicyType int32

	const (
		Policy_MIN PolicyType = 0
		Policy_MAX PolicyType = 1
		Policy_MID PolicyType = 2
		Policy_AVG PolicyType = 3
	)

	// 默认值 数字类型是0，字符串类型是空字符串，byte是0（本质是int8）
	var ok map[string]int
	fmt.Println(ok)

	// iota 特殊常量，可以认为是一个可以被编译器修改的常量
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	MyMap := map[string]interface{}{
		"1": 1,
		"2": "string",
		"3": 12.12,
	}

	for _, value := range MyMap {
		if _, ok := value.(int); ok {
			continue
		}
		switch v := value.(type) {
		case *string:
			fmt.Printf("Type *string %v\n", v)
		case string:
			fmt.Printf("Type string %v\n", v)
		case float32:
			fmt.Printf("Type float32 %v\n", v)
		case nil:
			fmt.Println("nil value: nothing to check?")
		default:
			fmt.Printf("Unexpected type %v\n", v)
		}
	}

	var strInt = int16(-255)
	dstInt := int8(strInt)
	fmt.Println(strInt)
	fmt.Println(dstInt)

	fmt.Println("----------------")

	str := "Hello!"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	fmt.Println("=============")
	for _, char := range str {
		fmt.Printf("%c", char)
	}

}
