package interface_test

import "testing"

type Programmer interface {
	WriteHelloWordCode() string
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWordCode() string{
	return "fmt.Println("+"Hello World!"+")"
}


func TestInterface(t *testing.T){
	var p Programmer = new(GoProgrammer)
	t.Logf("%T",p)
	t.Log(p.WriteHelloWordCode())
}