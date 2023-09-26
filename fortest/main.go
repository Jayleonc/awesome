package fortest

import "fmt"

func main() {
	// 1. 打印 0-50 所有的偶数
	/*	for i := 0; i <= 50; i++ {
			if i%2 == 0 {
				println("i = ", i)
			}
		}
	*/
	// 2. 求 1 到 100 所有数相加的和
	/*	sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		println("sum = ", sum)
	*/
	// 3. 打印1 ～ 100 之间所有是 9 的倍数的整数的个数及总和
	/*	var sum = 0
		var count = 0
		for i := 1; i <= 100; i++ {
			if i%9 == 0 {
				sum += i
				count++
			}
		}
		println("sum = ", sum, " count = ", count)
	*/

	// 4. 计算5的阶乘 （1*2*3*4*5）
	/*	var sum = 1
		for i := 1; i <= 5; i++ {
			sum *= i
		}
		println("sum = ", sum)
	*/

	// 5. 打印一个矩形
	/*	for i := 1; i <= 16; i++ {
			print("*")
			if i%4 == 0 {
				println("")
			}
		}
	*/

	// 6. for 循环的嵌套
	/*	var row = 5
		var col = 4
		for i := 0; i < row; i++ {
			for o := 0; o < col; o++ {
				print("*")
			}
			println("")
		}
	*/

	var arr = []string{"php", "java", "c++", "golang"}

	for _, val := range arr {
		fmt.Println(val)
	}

}
