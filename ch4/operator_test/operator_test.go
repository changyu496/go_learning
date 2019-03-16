package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executeble
)

// Go语言中数组的维数、个数相等时，可以比较，否则编译无法通过
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{5, 6, 7, 8}
	c := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c)
}

// 按位清空，Go语言特有
func TestBit(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeble == Executeble)
}
