package main

import (
	"fmt"
	"reflect"
)

func Add(a, b, c int) int {
	return a + b + c
}

func main() {
	v := reflect.ValueOf(Add) // 获得 函数 Add 对应的反射对象
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn()) // t.NumIn 获取函数的入参个数
	fmt.Println(len(argv))
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i) // 设置 argv 数组的各个参数
	}
	argv2 := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3)}
	result := v.Call(argv2) // 传入参数列表
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // 返回数组长度
}
