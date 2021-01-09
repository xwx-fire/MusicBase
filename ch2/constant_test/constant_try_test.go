package constant_test

import "testing"

const (
	Monday = iota+1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const(
	open = 1 << iota
	read
	write
	close
)

func TestConstantTry(t *testing.T){
	t.Log(Monday,Thursday)
	t.Log(open,read,write,close)
}

func TestConst(t *testing.T){
	a := 7 //0111
	t.Log(a&open == open,a&read == read, a&write == write, a&close == close)
}