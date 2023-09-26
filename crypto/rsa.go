package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func main() {
	prvKey, pubKey := GenRsaKey()
	text := []string{"天青色等烟雨，而我在等你，", "无关风月，我题序等你回"}
	var context []byte
	for i := 0; i < len(text); i++ {
		context = append(context, text[i]...)
	}
	ciphertext := RsaEncrypt(context, pubKey)
	fmt.Println("密文为：", ciphertext)
	result := RsaDecrypt(ciphertext, prvKey)
	fmt.Println("明文为：", string(result))

	// 签名
	block, _ := pem.Decode(prvKey)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	// 签名值
	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.Hash(0), context)
	if err != nil {
		return
	}

	block, _ = pem.Decode(pubKey)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := publicKey.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.Hash(0), context, sign)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("验证成功")
	}

}

func RsaEncrypt(data, keyBytes []byte) []byte {
	// 解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}

// RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}
