package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s1 := "A,B,C"
	parts := strings.Split(s1, ",")
	for _, v := range parts {
		t.Log(v)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Logf("%s", s)
	if v, err := strconv.Atoi("10"); err == nil {
		t.Log(v)
	}
}
