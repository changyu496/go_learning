package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Print(" ", host)
}

type Dog struct {
	//p *Pet
	Pet
}

// func (d *Dog) Speak() {
// 	fmt.Print("Wang")
// }

// func (d *Dog) SpeakTo(host string) {
// 	d.p.SpeakTo(host)
// }

func TestDog(t *testing.T) {
	// Go 不支持强制类型转换
	dog := new(Dog)
	dog.SpeakTo("Yu")
}
