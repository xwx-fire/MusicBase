package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T){
	counter := 0
	i := 0
	for i < 5000{
		go func() {
			counter++
		}()
		i++
	}
	time.Sleep(time.Second*1)
	t.Logf("counter = %d",counter)
}


func TestCounterThreadSafe(t *testing.T){
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++{
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second*1)
	t.Logf("counter = %d",counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}