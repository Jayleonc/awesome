package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 创建自定义的 TLS 配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过验证服务器的证书有效性（仅用于测试，生产环境中不应使用）
	}

	// 创建带有自定义 TLS 配置的 Transport
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// 创建一个使用自定义 Transport 的 HTTP 客户端
	client := &http.Client{
		Transport: transport,
	}

	// 发起 HTTPS 请求
	response, err := client.Get("https://localhost:8443/") // 替换为你的实际请求 URL
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}

	fmt.Println("响应体:", string(body))
}
