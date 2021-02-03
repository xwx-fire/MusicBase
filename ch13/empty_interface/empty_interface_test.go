package empty_interface

import (
	"fmt"
	"testing"
)

/**
	空接口测试
 */

func DoSomething(p interface{}){
	if v,ok := p.(int);ok{
		fmt.Println("int",v)
	}
	if v,ok := p.(string);ok{
		fmt.Println("string",v)
	}

	switch v:= p.(type){
	case int:
		fmt.Println("int",v)
	case string:
		fmt.Println("string",v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T){
	//DoSomething(10)
	//DoSomething("10")
	//var p interface{}
	//p = []int{1,2,3}
	//t.Logf("%T",p)
	//t.Log(reflect.TypeOf(p))
	var a float64 = 1800
	var b float64 = 1799.99
	t.Log(a - b)
}

func CeilFloat(a,b float64){

}


