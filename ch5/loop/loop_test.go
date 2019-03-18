package loop_test

import "testing"

/*Go语言的就一个循环关键字，for，玩法和while一致*/
func TestWhile(t *testing.T) {
	n := 1
	for n < 5 {
		t.Log(n)
		n++
	}
}
