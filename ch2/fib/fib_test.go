package fib

import (
	"testing"
)

func TestFibList(t *testing.T){
	//var a int = 1
	//var b int = 1


	var(
		a = 1
		b = 1
	)

	t.Log(a)
	for i := 0; i < 5; i++{
		t.Log(b)
		fmp := a
		a = b
		b = fmp+a
	}
	t.Log()
}

func TestExChange(t *testing.T){
	a := 1
	b := 2
	a,b = b,a
	t.Log(a,b)
}


