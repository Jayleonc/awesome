package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateCSR() {
	// 生成 RSA 密钥对
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("生成 RSA 密钥对失败:", err)
		os.Exit(1)
	}

	// 定义 CSR 中的 Subject 信息
	subject := pkix.Name{
		CommonName:         "www.example.com",
		Organization:       []string{"Example Inc."},
		OrganizationalUnit: []string{"IT"},
		Locality:           []string{"San Francisco"},
		Province:           []string{"CA"},
		Country:            []string{"US"},
	}

	// 创建 CSR 请求
	csrTemplate := x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: x509.SHA256WithRSA, // 用于指定 用于生成证书请求的私钥 的签名算法。在 X.509 标准中，私钥用于对 证书请求 (CSR就是证书请求)// 进行签名，生成一个数字签名，以验证证书请求的真实性和完整性。
	}

	// rand: 一个随机数生成器，用于生成证书请求中的随机数。可以使用 crypto/rand 包中的函数生成。
	// TODO: rand 的随机数生成器是怎么实现的
	// template: 一个 CertificateRequest 结构体，包含了证书请求的模板信息，如公钥、标识信息等。
	// priv: 一个私钥，用于对证书请求进行签名。
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, key)
	if err != nil {
		fmt.Println("创建 CSR 请求失败:", err)
		os.Exit(1)
	}

	// 将 CSR 请求保存到文件
	csrFile, err := os.Create("csr.pem")
	if err != nil {
		fmt.Println("创建 CSR 文件失败:", err)
		os.Exit(1)
	}
	defer csrFile.Close()

	err = pem.Encode(csrFile, &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	})
	if err != nil {
		fmt.Println("保存 CSR 文件失败:", err)
		os.Exit(1)
	}

	fmt.Println("成功生成 CSR 文件: csr.pem")
}
