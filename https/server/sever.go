package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, HTTPS!")
	})

	// 设置 TLS 配置
	tlsConfig := &tls.Config{
		// 生成自签名证书和私钥，仅用于示例目的
		Certificates: []tls.Certificate{},
	}

	// 创建 HTTPS 服务器
	server := &http.Server{
		Addr:      ":8443", // 监听地址和端口
		TLSConfig: tlsConfig,
	}

	log.Println("启动 HTTPS 服务器...")
	log.Fatal(server.ListenAndServeTLS("https/server.crt", "https/server.key"))
}
