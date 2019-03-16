package fib

import "testing"

// 斐波那契数列的实现
func TestFibList(t *testing.T) {
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		temp := b
		b = a + b
		a = temp
	}
}

// 交换两个变量
func TestExchange(t *testing.T) {
	a := "a"
	b := "b"
	a, b = b, a
	t.Log(a, b)
}
