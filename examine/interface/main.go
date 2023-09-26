package main

import "fmt"

type reader interface {
	read1(b []byte) (int, error)

	read2(n int) ([]byte, error)
}

type tom struct {
	name string
}

func main() {
	var t tom
	var r reader
	r = t
	str := []byte("hello")
	r.read1(str)
	r.read2(1)
	ch := make(chan int, 8193)
	ch <- 10
}

func (u tom) read1(b []byte) (int, error) {
	fmt.Println(b)
	return 0, nil
}

func (u tom) read2(n int) ([]byte, error) {
	s := make([]byte, 1024)
	return s, nil
}
