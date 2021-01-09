package operator_test

import "testing"

const(
	open = 1 << iota
	read
	write
	close
)

func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3,4}
	b := [...]int{5,6,7,8}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}

	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T){
	a := 7 //0111
	a = a &^ read
	t.Log(a)
	t.Log(a&open == open,a&read == read, a&write == write, a&close == close)
}