package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnRandValue()(int,int){
	return rand.Intn(10),rand.Intn(20)
}

func timeSpend(inner func(int) int) func(int) int{
	return func(p int) int {
		se := time.Now()
		res := inner(p)
		fmt.Println("time spend",time.Since(se).Minutes())
		return res
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T){
	t.Log(returnRandValue())
	tsF := timeSpend(slowFun)
	t.Log(tsF(10))
}

func Sum(op...int) int {
	ret := 0
	for _,v := range op {
		ret += v
	}
	return ret
}

func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear(){
	fmt.Println("clear resources")
}

func TestDefer(t *testing.T){
	defer Clear()
	t.Log("start")
	panic("error")
}