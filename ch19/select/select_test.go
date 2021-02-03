package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string{
	time.Sleep(time.Millisecond*500)
	return "Done"
}

func AsyncService() chan string{
	retch := make(chan string)
	go func(){
		ret := service()
		fmt.Println("return result")
		retch <- ret
		fmt.Println("service exited")
	}()
	return  retch
}

func TestSelect(t *testing.T){
	select {
		case ret := <- AsyncService():
			t.Log(ret)
		case <- time.After(time.Millisecond*100):
			t.Error("time out")
	}
}