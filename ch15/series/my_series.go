package series

import "fmt"

/**
	init函数先执行，可以有多个init
 */

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

/**
	大写字母开头的变量或方法才会被其他的包可见
 */
func Square(n int) int {
	return n * n
}

func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}