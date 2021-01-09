package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/**
	自定义error变量以便后续检索
 */
var LessThanTwoError = errors.New("n should larger than 2")
var LargerThanHundredError = errors.New("n should less than 100")

func GetFibonacci(n int) ([]int, error){
	if n < 2 {
		return nil,LessThanTwoError
	}
	if n > 100 {
		return nil,LargerThanHundredError
	}
	fibList := []int{1,1}
	for i := 2; i < n; i++{
		fibList = append(fibList,fibList[i - 2]+fibList[i - 1])
	}
	return fibList,nil
}

func TestGetFibonacci(t *testing.T){
	if v,err := GetFibonacci(1);err != nil{
		t.Log(err)
	}else{
		t.Log(v)
	}
}


func StringConvToInt(str string){
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil{
		if list, err = GetFibonacci(i); err == nil{
			fmt.Println(list)
		}else{
			fmt.Println("Error",err)
		}
	}else{
		fmt.Println("Error",err)
	}
}

/**
	失败退出优先，避免嵌套
 */
func StringConvToInt2(str string){
	var (
		i int
		err error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil{
		fmt.Println("Error ",err)
		return
	}
	if list,err = GetFibonacci(i); err != nil{
		fmt.Println("Error ",err)
	}
	fmt.Println(list)
}
