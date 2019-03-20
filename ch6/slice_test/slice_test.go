package slice_test

import "testing"

/*Go语言中的切片是比较难理解的，这里是切片的初始化方法*/
func TestSliceInit(t *testing.T) {
	a1 := []int{}
	a2 := make([]int, 3, 5)
	t.Log(a1, len(a1), cap(a1))
	t.Log(a2, len(a2), cap(a2))
}

/*Go语言中切片的动态扩容*/
func TestSliceGrowing(t *testing.T) {
	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, 1)
		t.Log(a, len(a), cap(a))
	}
}

/*Go语言中的切片和共享内存*/
func TestSliceShareMemory(t *testing.T) {
	year := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	q2 := year[3:6]
	q3 := year[6:9]
	q2[0] = "unknown"
	t.Log(q2, len(q2), cap(q3))
	t.Log(q3, len(q3), cap(q3))
	t.Log(year, len(year), cap(year))
}
