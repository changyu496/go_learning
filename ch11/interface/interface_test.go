package interface_test

import "testing"

type Programmer interface {
	WriteCode() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteCode() string {
	return "fmt.Println('hello')"
}

/*Duck Type类型*/
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteCode())
}
