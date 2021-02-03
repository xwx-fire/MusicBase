package channel_close

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func dataProducer1(ch chan int,wg *sync.WaitGroup){
	go func() {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println("写入数据错误")
			}
		}()
		i := 0
		for i < 10{
			select {
			case ch <- i:
				fmt.Printf("已存入%d\n",i)
				//ch <- i
			default:
				//close(ch)
				fmt.Println("channel is full")
				//break l
			}
			//fmt.Printf("已存入%d\n",i) //当channel buffer 存满时线程阻塞等待buffer中出现可容纳新数据的位置
			//ch <- i
			i++
		}
		close(ch)
		wg.Done()
	}()
}

func dataProducer2(ch chan string,wg *sync.WaitGroup){
	str := "123456"
	go func() {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println("错误",err)
			}
		}()
		l:for i := 1; i < len(str); i++ {
			select {
				case ch <- str[i-1:i]:
					fmt.Printf("存入字符串%v\n",str[i-1:i])
				default:
					break l
			}
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver1(ch chan int,wg *sync.WaitGroup){
	go func() {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println("接受数据错误")
			}
		}()
		for {
			if data,ok := <- ch;ok{
				fmt.Println(data)
			}else{
				fmt.Println("未找到元素")
				break
			}
		}
		wg.Done()
	}()
}

func dataReceiver2(ch chan string,ch1 chan int,wg *sync.WaitGroup){
	go func() {
		l:for{
			select {
			case data, ok := <- ch:
				if ok {
					fmt.Printf("取出管道1中的字符串%s\n",data)
				}else{
					fmt.Println("未找到元素")
				}
			case data, ok := <-ch1:
				if ok {
					fmt.Printf("取出管道2中的整数%d\n",data)
				}else{
					fmt.Println("未找到该元素")
				}
			default :
				//fmt.Println(<-ch)
				//fmt.Println(<-ch)
				//fmt.Println(<-ch)
				//fmt.Println(<-ch)
				//fmt.Println(<-ch)
				fmt.Println("没有取到其中任何一个管道中的数据")
				break l
			}
			//if data,ok := <- ch;ok{
			//	fmt.Printf("取出管道1中的字符串%s\n",data)
			//}else if data1,ok := <- ch1;ok{
			//	fmt.Printf("取出管道2中的整数%d\n",data1)
			//}else{
			//	fmt.Println("没有取到其中任何一个管道中的数据")
			//	break l
			//}
		}
		wg.Done()
	}()
}

func TestMoreChannel(t *testing.T){
	var wg sync.WaitGroup
	//var wg2 sync.WaitGroup
	ch := make(chan int,5)
	ch2 := make(chan string,5)
	wg.Add(1)
	dataProducer1(ch,&wg)
	wg.Add(1)
	dataProducer2(ch2,&wg)
	//wg.Add(1)
	//dataReceiver1(ch,&wg)
	time.Sleep(time.Second*3)
	wg.Add(1)
	dataReceiver2(ch2,ch,&wg)
	wg.Wait()
}