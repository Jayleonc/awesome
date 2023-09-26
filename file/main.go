package main

import (
	"fmt"
)

func main() {
	// 读取文件路径
	path := "/Users/jayleonc/Downloads/lr预设/"
	//open, err := os.ReadDir(path)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println(path)

	{
		path := "adf"
		fmt.Println(path)
	}

	path = "ok"
	fmt.Println(path)

	//for _, i := range open {
	//	if i.IsDir() {
	//		dir, _ := os.ReadDir(path + i.Name() + "/")
	//		for _, i2 := range dir {
	//			if i2.IsDir() {
	//				dir2, _ := os.ReadDir(path + i.Name() + "/" + i2.Name())
	//				for _, i3 := range dir2 {
	//					//fmt.Println(path + i.Name() + "/" + i2.Name() + "/" + i3.Name())
	//					if strings.Contains(i3.Name(), "~~") {
	//						_, after, _ := strings.Cut(i3.Name(), "~~")
	//						os.Rename(path+i.Name()+"/"+i2.Name()+"/"+i3.Name(), path+i.Name()+"/"+i2.Name()+"/"+after)
	//					}
	//					if i3.IsDir() {
	//						dir3, _ := os.ReadDir(path + i.Name() + "/" + i2.Name() + "/" + i3.Name())
	//						for _, i4 := range dir3 {
	//							_, after, _ := strings.Cut(i4.Name(), "~~")
	//							os.Rename(path+i.Name()+"/"+i2.Name()+"/"+i3.Name()+"/"+i4.Name(), path+i.Name()+"/"+i2.Name()+"/"+i3.Name()+"/"+after)
	//						}
	//					}
	//				}
	//			}
	//		}
	//	}
	//}
}

func show() {
	path := "ok"
	fmt.Println(path)
}
