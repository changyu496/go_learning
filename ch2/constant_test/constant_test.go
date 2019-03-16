package constant_test

import "testing"

const (
	Mon = iota + 1
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func TestConstantTry(t *testing.T) {
	t.Log(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
}
