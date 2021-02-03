package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string{
	time.Sleep(time.Millisecond*50)
	return "Done"
}

func otherTask(){
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

//func TestService(t *testing.T){
//	fmt.Println(service())
//	otherTask()
//}

func AsyncService() chan string{
	retCh := make(chan string,2) //buffer channel 若buffer中能装入数据则发送方向里存数据，若buffer中数据为空接收方必须等待数据，程序可以继续向下执行
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()

	return retCh
}

func TestAsyncService(t *testing.T){
	retCh := AsyncService()
	//t.Log(cap(retCh))
	//defer func(ret chan string) {
	//	close(ret)
	//}(retCh)
	otherTask()
	fmt.Println(<- retCh)
	//time.Sleep(time.Second*1)
}