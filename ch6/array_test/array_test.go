package array_test

import "testing"

/*Go语言初始化一个数组*/
func TestArrayInit(t *testing.T) {
	var a1 [3]int
	t.Log(a1)
	a2 := [4]int{1, 2, 3, 4}
	t.Log(a2)
	a3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	t.Log(a3)
}

/*Go语言的数组循环的方式*/
func TestArrayTravel(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		t.Log(a[i])
	}
	for idx, e := range a {
		t.Log(idx, e)
	}
	for _, e := range a {
		t.Log(e)
	}
}

/*Go语言的数组截取，很类似于Python*/
func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	a_section := a[1:4]
	t.Log(a_section)
}
