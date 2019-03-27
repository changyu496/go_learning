package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   int
	Name string
	Age  int
}

// func (e *Employee) ToString() string {
// 	fmt.Printf("ID:%d-Name:%s Age:%d", e.ID, e.Name, e.Age)
// }

func (e *Employee) String() string {
	fmt.Printf("Address:%x\r\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%d-Name:%s Age:%d", e.ID, e.Name, e.Age)
}

/*三种创建对象的方法，指针和对象，一般推荐后面两种方法*/
func TestCreateEmployObj(t *testing.T) {
	e := Employee{1, "ChangYu", 28}
	e1 := Employee{Name: "ChangYu", Age: 28}
	e2 := new(Employee)
	e2.Age = 28
	e2.Name = "ChangYu"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func TestOperateObjFunc(t *testing.T) {
	e := Employee{1, "ChangYu", 28}
	fmt.Printf("Address:%x\r\n", &e.Name)
	fmt.Print(e.String())
}
