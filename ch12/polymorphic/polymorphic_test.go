package polymorphic_test

import "testing"

type Code string
type Programmer interface {
	WriteFirstProgramme() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println('hello world')"
}

type JavaProgrammer struct {
}

func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println('hello world')"
}

func TestPolymorphic(t *testing.T) {
	progG := &GoProgrammer{}
	progJava := new(JavaProgrammer)
	t.Log(progG.WriteHelloWorld())
	t.Log(progJava.WriteHelloWorld())
}
