package map_test

import "testing"

func TestMapWithFunValue(t *testing.T){
	m1 := map[int]func(p int)int{}
	t.Log(m1)
	m1[1] = func(p int) int{return p}
	m1[2] = func(p int) int{return p*p}
	m1[3] = func(p int) int{return p*p*p}
	t.Log(m1)
	t.Log(m1[1](2),m1[2](2),m1[3](2))
}

/*
	map 实现set
	set 特点
	添加元素
	删除元素
	判断元素是否存在
	元素个数
*/
func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n]{
		t.Logf("%d is existing",n)
	}else{
		t.Logf("%d is not existing",n)
	}

	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet,1)
	n = 1
	if mySet[n]{
		t.Logf("%d is existing",n)
	}else{
		t.Logf("%d is not existing",n)
	}
}