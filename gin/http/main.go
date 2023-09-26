package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	var buf bytes.Buffer
	var l listen
	var s http.Server
	s.Addr = ":8080"
	s.Handler = l
	s.WriteTimeout = time.Second
	s.ErrorLog = log.New(&buf, "log:", log.LstdFlags)
	s.ConnState = l.ConnState
	err := s.ListenAndServe()
	if err != nil {
		s.ErrorLog.Print()
		fmt.Println(&buf)
	}
}

type listen struct {
	nu int
}

func (l listen) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/home" {
		w.Write([]byte("This is home"))
		return
	}
	if r.URL.Path == "/ping" {
		encode := r.URL.Query().Encode()
		w.Write([]byte(encode))
		return
	}
	w.Write([]byte("<h1>404</h1>"))
}

func (l listen) ConnState(coon net.Conn, c http.ConnState) {
	fmt.Println(c.String(), coon.LocalAddr())
}
