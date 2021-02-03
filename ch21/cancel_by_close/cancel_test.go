package cancel_by_close

import (
	"fmt"
	"sync"
	"testing"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
		case <- cancelChan:
			return true
		default:
			return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}

}

func cancel_2(caccelChan chan struct{}) {
	close(caccelChan)
}

func TestCancel(t *testing.T) {
	var wg sync.WaitGroup
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++{
		go func(i int, cancelCh chan struct{}) {
			wg.Add(1)
			for {
				if isCancelled(cancelCh){
					break
				}
			}
			fmt.Println(i,"Cancelled")
			wg.Done()
		}(i, cancelChan)
	}

	cancel_1(cancelChan)
	wg.Wait()
}