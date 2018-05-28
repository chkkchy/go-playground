package dbrw

import (
	"testing"
)

func TestDo(t *testing.T) {
	var rw Iface = ReaderWriter{
		Reader: Reader{Shared: "rrr"},
		Writer: Writer{Shared: "www"},
	}
	rw.Read()
	rw.Write()
}
