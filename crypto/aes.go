package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密函数
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	key2, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("hello ming")

	c := make([]byte, aes.BlockSize+len(plaintext))
	iv := c[:aes.BlockSize]

	// 加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		fmt.Println(err)
	}

	// 打印加密base64后密文
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	// 解密
	plaintext, err = AesDecrypt(ciphertext, key2, iv)
	if err != nil {
		fmt.Println(err)
	}

	// 打印解密明文
	fmt.Println(string(plaintext))

}
