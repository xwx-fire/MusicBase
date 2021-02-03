package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int,wg *sync.WaitGroup){
	go func() {
		for i := 0; i < 10; i++{
			ch <- i
			fmt.Printf("存入管道%d\n",i)
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup){
	go func() {
		//for i := 0; i < 10; i++{
		//	data := <- ch
		//	fmt.Println(data)
		//}
		for{
			if data,ok := <-ch;ok{
				fmt.Printf("取出数据%d\n",data)
			}else{
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T){
	var wg sync.WaitGroup
	ch := make(chan int,3)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Wait()
}