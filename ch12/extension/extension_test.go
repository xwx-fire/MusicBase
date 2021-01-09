package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak(){
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string){
	p.Speak()
	fmt.Print(" ",host)
}

type Dog struct {
	Pet
}

//func (d *Dog) Speak(){
//	fmt.Printf("Wang!")
//}
//
//func (d *Dog) SpeakTo(host string){
//	d.Speak()
//	fmt.Print(" ",host)
//}


func TestExtension(t *testing.T){
	dog := new(Dog)
	dog.SpeakTo("xxx")
}
