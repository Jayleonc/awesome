package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://www.baidu.com"

	// 根据URL获取资源
	res, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	// 读取资源数据 body: []byte
	body, err := ioutil.ReadAll(res.Body)

	// 关闭资源流
	res.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	// 控制台打印内容 以下两种方法等同
	fmt.Printf("%s", body)
	fmt.Printf(string(body))

	// 写入文件
	ioutil.WriteFile("site.txt", body, 0644)
}
