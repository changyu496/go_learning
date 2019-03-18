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
