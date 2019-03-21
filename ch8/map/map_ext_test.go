package map_test

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m1 := map[int]func(o int) int{}
	m1[1] = func(o int) int { return o }
	m1[2] = func(o int) int { return o * o }
	m1[3] = func(o int) int { return o * o * o }
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	delete(mySet, 1)
	t.Log(len(mySet))
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
