package main

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/tjfoc/gmsm/sm2"
	"os"
)

const (
	path = "awesome/crypto/key/"
)

func main() {
	privateKey := GetSM2PrivateKey()
	GetPCKS8PrivateKey(privateKey)
}

func GetPCKS8PrivateKey(privateKey *sm2.PrivateKey) *os.File {
	derBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil
	}
	file, err := os.OpenFile(path+"private_key.pkcs8", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	defer file.Close()

	if err = pem.Encode(file, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derBytes,
	}); err != nil {
		return nil
	}
	return file
}

// GetSM2PrivateKey 提供 SM2 私钥
func GetSM2PrivateKey() *sm2.PrivateKey {
	random := make([]byte, 16)
	reader := rand.Reader
	if _, err := reader.Read(random); err != nil {
		panic(err)
	}

	key, err := sm2.GenerateKey(reader)
	if err != nil {
		panic(err)
	}

	return key
}

// GetSM2PublicKey 提供 SM2 公钥
func GetSM2PublicKey(key *sm2.PrivateKey) sm2.PublicKey {
	return key.PublicKey
}
