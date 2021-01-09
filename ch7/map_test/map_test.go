package map_test

import "testing"

func TestMapInit(t *testing.T){
	m0 := map[int]int{}
	t.Log(m0,len(m0))

	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1)
	t.Logf("len m1 = %d",len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2)
	t.Log()
}

func TestAccessNotExistingKey(t *testing.T){
	m1 := map[int]int{}
	t.Log(len(m1))
	m1[2] = 0
	t.Log(len(m1))
	m1[3] = 0
	if v, ok := m1[3];ok{
		t.Logf("Key 3's value is %d",v)
	}else{
		t.Log("Key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T){
	m1 := map[int]int{1:1, 2:4, 3:9}
	for k, v := range m1{
		t.Log(k, v)
	}
}