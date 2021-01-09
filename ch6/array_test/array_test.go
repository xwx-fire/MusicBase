package array_test

import "testing"

func TestArrayInit(t *testing.T){
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,3,5,7}

	t.Log(arr[0], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T){
	arr3 := [...]int{1,2,3,4,5}
	for i := 0; i < len(arr3); i++{
		t.Log(arr3[i])
	}

	for _, e := range arr3{
		t.Log(e)
	}

	//arr3 := [...][2]int{{1,2},{3,4}}
	//t.Log(arr3)
	//for i := 0; i < len(arr3); i++{
	//	t.Log(arr3[i])
	//	for j := 0; j < len(arr3[i]); j++{
	//		t.Log(arr3[i][j])
	//	}
	//}
}

func TestArraySection(t *testing.T){
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}
