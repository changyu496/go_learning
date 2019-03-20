package map_test

import "testing"

/*Go语言中三种初始化map的方法*/
func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1["one"])
	m2 := map[string]int{}
	m2["1"] = 1
	t.Log(m2["1"])
	m3 := make(map[string]int, 10)
	t.Log(m3)
	t.Log(m3["1"])
}
func TestAccessNotExistKey(t *testing.T) {
	m1 := map[string]int{"1": 1, "2": 2, "3": 3}
	m1["4"] = 1
	t.Log(m1["4"])
	if v, ok := m1["4"]; ok {
		t.Logf("m1[4] value is %d", v)
	} else {
		t.Log("m1[4] value not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m1 {
		t.Logf("key:%s,value:%d", k, v)
	}
}
