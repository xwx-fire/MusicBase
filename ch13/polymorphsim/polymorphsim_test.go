package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() Code{
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {

}

func (J *JavaProgrammer) WriteHelloWorld() Code{
	return "System.Out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer){
	fmt.Printf("%T %v\n",p ,p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T){
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}