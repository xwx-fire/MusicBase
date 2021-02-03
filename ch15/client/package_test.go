package client

import (
	"ch15/series"
	"fmt"
	"testing"
)


func init(){
	fmt.Println("init3")
}

/*
	包名不一定要与目录名称一致
	同一目录下两个包名会冲突
 */
func TestPackage(t *testing.T){
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Square(5))
}

