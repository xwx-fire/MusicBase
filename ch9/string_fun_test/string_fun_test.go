package string_fun_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T){
	s := "A,B,C"
	parts := strings.Split(s,",")
	t.Log(parts[0],parts[1],parts[2])
	t.Log(strings.Join(parts,"-"))
}

func TestStringConv(t *testing.T){
	s := strconv.Itoa(10)
	t.Log("str"+s)
	t.Log(s)

	if v,err := strconv.Atoi("10");err == nil{
		t.Log(10+v)
	}
}