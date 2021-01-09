package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id int
	Name string
	Age int
}

func (e Employee) String() string{
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e))
	return fmt.Sprintf("ID:%d-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

//func (e *Employee) String() string{
//	fmt.Printf("Address is %x\n",unsafe.Pointer(e))
//	return fmt.Sprintf("ID:%d/Name:%s/Age:%d",e.Id,e.Name,e.Age)
//}

func TestCreateEmployeeObj(t *testing.T){
	e := Employee{0,"Bob",20}
	t.Log(e)
	e1 := Employee{Name:"Mike",Age:23}
	t.Log(e1)
	e2 := new(Employee)
	t.Log(e2)

	t.Log("初始化之后")
	e2.Id = 2
	e2.Name = "Joe"
	e2.Age = 24
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T",e)
	t.Logf("e2 is %T",e2)
}

func TestStructOperations(t *testing.T){
	e := Employee{0,"谢文轩",20}
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e))
	t.Log(e.String())
}