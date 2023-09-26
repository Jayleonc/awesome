package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func loadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal("ReadFile: ", err)
	}
	pool.AppendCertsFromPEM(ca)
	return pool
}

func main() {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: loadCA("/Users/jayleonc/server.crt")},
		}}

	resp, err := c.Get("https://localhost:5200")
	if err != nil {
		log.Fatal("http.Client.Get: ", err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
