package ncpu_test

import (
	"fmt"
	"sync"
	"testing"
)

type Vector []float64

const NCPU = 3

func (v *Vector)DoSome(u Vector, ch chan float64,wg *sync.WaitGroup){
	sum := 0.00
	fmt.Println(u)
	for _,v1 := range u{
		sum += v1
	}

	ch <- sum
	wg.Done()
}

func (v *Vector)DoAll(wg *sync.WaitGroup){
	ch := make(chan float64,NCPU)
	for i := 0; i < NCPU; i++{
		t := (*v)[i*len(*v)/NCPU :(i+1)*len(*v)/NCPU]
		wg.Add(1)
		go v.DoSome(t,ch,wg)
	}
	//i := 0
	//	wg.Add(1)
	//	l:for {
	//		select {
	//		case <- ch:
	//			fmt.Printf("第%d片CPU完成计算\n",i)
	//		default:
	//			fmt.Println("获取失败")
	//			break l
	//		}
	//		i++
	//	}
	//	wg.Done()
	for i := 0;i < NCPU;i++{
		if data,ok := <- ch;ok{
			fmt.Printf("第%d片CPU完成计算\n",i)
			fmt.Println(data)
		}else{
			fmt.Println("获取失败")
		}
	}
	close(ch)
}

func TestNcpu(t *testing.T){
	var wg sync.WaitGroup
	v := Vector{1,2,3,4,5,6,7,8,9,10}
	v.DoAll(&wg)
	wg.Wait()
	t.Log(v)
}
