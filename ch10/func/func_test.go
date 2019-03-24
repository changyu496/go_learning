package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(100), rand.Intn(20)
}

func TestReturnMultiValues(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slow(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsFn := timeSpent(slow)
	tsFn(1)
}
