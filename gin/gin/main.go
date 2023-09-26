package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func main() {
	c := New()
	c.Run()
}

type RW interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.Pusher
	http.CloseNotifier
	Pusher() http.Pusher
}

type rw struct {
	http.ResponseWriter
	size       int
	statuscode int
}

func (r *rw) Push(target string, opts *http.PushOptions) error {
	//TODO implement me
	panic("implement me")
}

func (r *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func (r *rw) CloseNotify() <-chan bool {
	return r.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (r *rw) Flush() {
	r.ResponseWriter.(http.Flusher).Flush()
}

func (r *rw) Pusher() (pusher http.Pusher) {
	if pusher, ok := r.ResponseWriter.(http.Pusher); ok {
		return pusher
	}
	return nil
}

type Context struct {
	Request *http.Request
	writer  rw
}

func (c *Context) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Request = r
	c.writer.ResponseWriter = w
	c.writer.ResponseWriter.Header().Set("content-type", "text/html;charset=utf-8")
	c.writer.ResponseWriter.Write([]byte("Hello Jay"))
}

var _ RW = &rw{}
var _ http.ResponseWriter = &rw{}
var _ http.Handler = &Context{}

func (c *Context) Run() (err error) {
	fmt.Println("启动成功")
	err = http.ListenAndServe(":8080", c)
	return
}

func New() *Context {
	return &Context{}
}
