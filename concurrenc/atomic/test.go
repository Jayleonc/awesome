package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int32 = 10
	swapped := atomic.CompareAndSwapInt32(&i, 10, 12)
	fmt.Printf("%v, %d\n", swapped, i)
	// 取地址里的值
	val := atomic.LoadInt32(&i)
	fmt.Printf("%v\n", val)
	// 往地址里存值
	atomic.StoreInt32(&i, 20)
	fmt.Printf("%d\n", i)

}
