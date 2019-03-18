package condition_test

import "testing"

/*Go语言的条件语句支持语句*/
func TestMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("do it")
	} else {
		t.Log("No!!!")
	}
}

/*Go语言的switch可以支持多个值*/
func TestSwitchMultiSec(t *testing.T) {
	a := 1
	for a < 5 {
		switch a {
		case 0, 2:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("不属于0到3")
		}
		a++
	}
}

/*Go语言的switch/case相当于if/else的用法*/
func TestSwitchCaseMultiSec(t *testing.T) {
	for a := 0; a < 5; a++ {
		switch {
		case a%2 == 0:
			t.Log("偶数")
		case a%2 == 1:
			t.Log("奇数")
		default:
			t.Log("奇怪的数")
		}
	}
}
