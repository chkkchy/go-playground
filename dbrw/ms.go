package dbrw

import "fmt"

type Iface interface {
	Read()
	Write()
}

type Reader struct {
	Shared string
}

func (r Reader) Read() {
	fmt.Println("reader: ", r.Shared)
}

type Writer struct {
	Shared string
}

func (r Writer) Write() {
	fmt.Println("writer: ", r.Shared)
}

type ReaderWriter struct {
	Reader
	Writer
}
