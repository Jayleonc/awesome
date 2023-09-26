package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 加载服务器证书和私钥
	// 它可以从文件或字节切片中加载 PEM 格式的 X.509 证书和私钥，并返回一个 tls.Certificate 对象，其中包含了加载的证书和私钥信息。
	cert, err := tls.LoadX509KeyPair("https/server.crt", "https/server.key")
	if err != nil {
		fmt.Println("加载服务器证书和私钥失败:", err)
		return
	}

	// 创建自定义的根 CA 证书池
	caPool := loadRootCA()
	if caPool == nil {
		return
	}

	// 创建自定义的客户端 CA 证书池
	clientCAPool := loadClientCA()
	if clientCAPool == nil {
		return
	}

	// 配置 TLS 客户端认证模式
	clientAuth := tls.RequireAndVerifyClientCert

	// 配置 TLS 版本、加密套件、椭圆曲线和会话票据
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert}, // 设置服务器证书和私钥
		ClientCAs:    clientCAPool,            // 设置客户端 CA 证书池
		ClientAuth:   clientAuth,              // 设置客户端认证模式
		RootCAs:      caPool,                  // 设置根 CA 证书池
		MinVersion:   tls.VersionTLS12,        // 设置支持的最小 TLS 版本
		CipherSuites: []uint16{ // 设置加密套件列表
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},
		CurvePreferences: []tls.CurveID{ // 设置椭圆曲线列表
			tls.CurveP256,
			tls.CurveP384,
			tls.CurveP521,
		},
		SessionTicketsDisabled: false, // 启用会话票据
	}

	// 创建自定义的 HTTP 传输
	// 用于控制 HTTP 请求的传输参数的结构体，它定义了诸如连接池、超时、TLS 配置等参数
	// 用于配置客户端在进行 HTTP 请求时的传输行为
	// 将其作为 http.Client 的 Transport 字段进行配置，可以实现对 HTTP 请求的传输参数进行定制化的配置
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// 创建自定义的 HTTP 客户端
	client := &http.Client{
		Transport: transport,
	}

	// 使用自定义的 HTTP 客户端进行 HTTPS 请求
	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("HTTPS 请求失败:", err)
		return
	}

	defer resp.Body.Close()

	// 处理响应
	fmt.Println("响应状态码:", resp.StatusCode)
	// ...
}

// 加载客户端 CA 证书
func loadClientCA() *x509.CertPool {
	caCert, err := ioutil.ReadFile("client-ca.crt")
	if err != nil {
		fmt.Println("加载客户端 CA 证书失败:", err)
		return nil
	}
	clientCAPool := x509.NewCertPool()
	clientCAPool.AppendCertsFromPEM(caCert)

	return clientCAPool
}

// 加载根 CA 证书
func loadRootCA() *x509.CertPool {
	caCert, err := ioutil.ReadFile("root-ca.crt")
	if err != nil {
		fmt.Println("加载根 CA 证书失败:", err)
		return nil
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	return caPool
}

// 以上示例代码演示了如何使用 Go 语言中的 `crypto/tls` 包配置自定义的 `tls.Config`，包括加载服务器证书和私钥、设置支持的 TLS 版本、加密套件、椭圆曲线、会话票据等，并设置客户端认证模式和加载客户端 CA 和根 CA 证书。同时，示例代码还包含了加载客户端 CA 和根 CA 证书的函数，并添加了注释以供参考。在实际应用中，你需要根据自己的实际情况替换证书文件的路径和文件名，并根据需要调整其他配置项。
