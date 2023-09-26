package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "011004"
	//funcName(str)
	funcName2(str)
	//calculationDigit(123456789)
}

func funcName2(str string) string {
	s1 := str[:len(str)-3]
	s2 := str[len(str)-3:]
	s2num, _ := strconv.Atoi(s2)
	s2dn := s2num + 1
	s2dnd := calculationDigit(s2dn)
	fill := fillZero(s2dnd, s2dn) // 对最后三位已经 +1 的进行补 0
	result := s1 + fill
	return result
}

func funcName(str string) {
	i2, _ := strconv.Atoi(str)
	strLen := calculationDigit(i2)
	if strLen == 17 || strLen == 16 {
		s1 := str[0:3]
		s2 := str[3:6] // 不准动这个
		s3 := str[6:9]
		s4 := str[9:12]
		s5 := str[12:15]
		s6 := str[15:18]
		s1num, _ := strconv.Atoi(s1)
		s1d := calculationDigit(s1num)
		fill0 := fillZero(s1d, s1num) // 对前三位进行补 0

		s6num, _ := strconv.Atoi(s6)
		s6dn := s6num + 1
		s6dnd := calculationDigit(s6dn)
		fill2 := fillZero(s6dnd, s6dn) // 对最后三位已经 +1 的进行补 0
		result := fill0 + s2 + s3 + s4 + s5 + fill2
		fmt.Println(result)
	} else if strLen == 14 || strLen == 13 {
		s1 := str[0:3]
		s2 := str[3:6] // 不准动这个
		s3 := str[6:9]
		s4 := str[9:12]
		s5 := str[12:15]
		s1num, _ := strconv.Atoi(s1)
		s1d := calculationDigit(s1num)
		fill0 := fillZero(s1d, s1num) // 对前三位进行补 0

		s5num, _ := strconv.Atoi(s5)
		s5dn := s5num + 1
		s5dnd := calculationDigit(s5dn)
		fill2 := fillZero(s5dnd, s5dn) // 对最后三位已经 +1 的进行补 0
		result := fill0 + s2 + s3 + s4 + fill2
		fmt.Println(result)
	} else if strLen == 11 || strLen == 10 {
		s1 := str[0:3]
		s2 := str[3:6] // 不准动这个
		s3 := str[6:9]
		s4 := str[9:12]
		s1num, _ := strconv.Atoi(s1)
		s1d := calculationDigit(s1num)
		fill0 := fillZero(s1d, s1num) // 对前三位进行补 0

		s4num, _ := strconv.Atoi(s4)
		s3dn := s4num + 1
		s3dnd := calculationDigit(s3dn)
		fill2 := fillZero(s3dnd, s3dn) // 对最后三位已经 +1 的进行补 0
		result := fill0 + s2 + s3 + fill2
		fmt.Println(result)
	} else if strLen == 8 || strLen == 7 {
		s1 := str[0:3]
		s2 := str[3:6] // 不准动这个
		s3 := str[6:9]
		s1num, _ := strconv.Atoi(s1)
		s1d := calculationDigit(s1num)
		fill0 := fillZero(s1d, s1num) // 对前三位进行补 0

		s3num, _ := strconv.Atoi(s3)
		s3dn := s3num + 1
		s3dnd := calculationDigit(s3dn)
		fill2 := fillZero(s3dnd, s3dn) // 对最后三位已经 +1 的进行补 0
		result := fill0 + s2 + fill2
		fmt.Println(result)
	} else if strLen == 4 || strLen == 5 {
		s1 := str[0:3]
		s2 := str[3:6]
		s1num, _ := strconv.Atoi(s1)
		s1d := calculationDigit(s1num)
		fill0 := fillZero(s1d, s1num) // 对前三位进行补 0

		s2num, _ := strconv.Atoi(s2)
		s2dn := s2num + 1
		s2dnd := calculationDigit(s2dn)
		fill2 := fillZero(s2dnd, s2dn) // 对最后三位已经 +1 的进行补 0
		result := fill0 + fill2
		fmt.Println(result)
	}
}

func fillZero(i int, atoi int) string {
	var fill0 string
	if i == 1 {
		fill0 = "00" + strconv.Itoa(atoi)
	} else if i == 2 {
		fill0 = "0" + strconv.Itoa(atoi)
	} else {
		fill0 = strconv.Itoa(atoi)
	}
	return fill0
}

func calculationDigit(n int) int {
	sum := 0
	for n != 0 {
		n /= 10
		sum++
	}
	fmt.Println("sum:", sum)
	return sum
}
