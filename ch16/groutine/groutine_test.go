package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGrouine(t *testing.T){
	for i := 0; i < 10; i++{
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond*2000)
}