package constant_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstanTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Monday + Tuesday)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
	t.Log(Readable, Writeable, Executable)
}
