package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

/**
	recover 获取panic中返回的Error
 */
func TestPanicVxExit(t *testing.T){
	defer func() {
		if err := recover();err != nil{
			fmt.Println("recovered from",err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something Wrong!"))
}