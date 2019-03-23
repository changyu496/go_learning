package string_test

import "testing"

func TestString(t *testing.T) {
	var s1 string = "hello"
	var s2 string
	s2 = "world"
	t.Log(s1)
	t.Log(s2)
	t.Logf("%s %s", s1, s2)

	s3 := "\xE4\xB8\xA5"
	t.Log(s3)

	s4 := "中"
	t.Log(s4)
	t.Log(len(s4))

	c := []rune(s4)
	t.Log(len(c))
	t.Logf("unicode:%x", c[0])
	t.Logf("utf8:%x", s4)

	for _, v := range c {
		t.Log(v)
	}
}
