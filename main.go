package main

import "github.com/chkkchy/go-playground/dbrw"

func main() {

	//var dbr dbrw.Iface = dbrw.Reader{}
	//dbr.Read()

	//var dbw dbrw.Iface = dbrw.Writer{}
	//dbw.Write()

	var rw dbrw.Iface = dbrw.ReaderWriter{
		Reader: dbrw.Reader{Shared: "rrr"},
		Writer: dbrw.Writer{Shared: "www"},
	}
	rw.Read()
	rw.Write()
}
