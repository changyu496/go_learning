package type_test

import "testing"

type MyInt int64

// 只能显示转换，不支持隐是转换
func TestImplict(t *testing.T) {
	var a int = 99
	var b int64 = 1
	a = int(b)
	t.Log(a, b)
}

// 指针不支持运算
func TestPtr(t *testing.T) {
	i := 1
	iPtr := &i
	t.Log(i, iPtr)
	t.Logf("%T %T", i, iPtr)
}

// 字符串默认是空字符串
func TestString(t *testing.T) {
	var str string
	t.Log("*" + str + "*")
	t.Log(len(str))
}
