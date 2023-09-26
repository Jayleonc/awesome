package main

import "fmt"

type reader interface {
	read(b []byte) (int, error)
}

type reader2 interface {
	reader
}

type file struct {
	name    string
	context []byte
}

type pipe struct {
	name string
}

func main() {
	str := []byte("hello")
	var f file
	f.context = str

	var p pipe
	retrieve(f)
	retrieve(p)

	var r reader
	var r2 reader2
	r = r2
	r = f
	f1 := r.(file)
	r2 = f1

	fmt.Println("ok")
}

func (f file) read(b []byte) (int, error) {
	fmt.Printf("%c", b)

	return 0, nil
}

func (f pipe) read(b []byte) (int, error) {
	return len(b), nil
}

// 多态来啦
func retrieve(r reader) error {

	data := make([]byte, 100)

	_, err := r.read(data)
	if err != nil {
		return err
	}

	return nil
}
