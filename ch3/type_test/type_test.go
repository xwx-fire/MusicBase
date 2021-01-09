package type_test

import "testing"

type Myint int64

func TestImplicit(t *testing.T){
	var a int = 1
	var b int64
	var c Myint
	b = int64(a)
	c = Myint(a)
	t.Log(a,b,c)
}

func TestPointer(t *testing.T){
	a := 1
	prt := &a
	t.Log(a,prt)
	t.Logf("%T %T",a,prt)
}