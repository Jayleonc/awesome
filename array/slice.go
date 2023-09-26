package main

import "fmt"

func main() {
	//myArr := []int{1, 2, 3, 4} // 动态数组 切片 slice
	//for index, value := range myArr {
	//	fmt.Println("index = ", index, ", value= ", value)
	//}
	//
	//var strSlice = []string{"java", "php", "golang", "c++"}
	//for _, v := range strSlice {
	//	fmt.Printf("str:%v\n", v)
	//}

	// 创建一个字符串切片
	// 其长度和容量都是 5 个元素
	// 或分别指定长度和容量
	// 不允许创建容量小于长度的切片
	//slice := make([]int, 3, 5)

	// 通过切片字面量来声明切片
	//slice := []int{10, 20, 30}

	// 使用索引生命切片
	//slice := []string{99: ""}

	//声明时不做初始化，就会创建一个 nil 切片
	//var slice []int
	//slice := make([]int, 0)
	//slice := []int{}

	//slice := []int{10, 20, 30, 40, 50}
	//
	//newSlice := append(slice, 60)
	//
	//fmt.Printf("%v -- %T\n", slice, slice)
	//fmt.Printf("%v -- %T\n", newSlice, newSlice)

	// 这是一个切片
	slice := make([]int, 4)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
		fmt.Printf("Index: %d value: %d\n", i, slice[i])
	}

	arr := make([]byte, 4)
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
	fmt.Println(arr[0])

	// 这是一个数组
	//var mychars [26]byte
	//
	//for i := 0; i < 26; i++ {
	//	mychars[i] = 'A' + byte(i)
	//	fmt.Printf("%c ", mychars[i])
	//}
	//fmt.Println()
	//
	//var myMaps map[int]string
	//myMaps = make(map[int]string, 10)
	//myMaps[0] = "宋江"
	//myMaps[1] = "武松"
	//myMaps[2] = "李逵"
	//myMaps[1] = "武大郎"
	//myMaps[3] = "李广"
	//myMaps[2] = "吴用"
	//fmt.Println(myMaps)
	//
}
