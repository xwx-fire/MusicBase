package condition

import "testing"

func TestCondition(t *testing.T){
	if a :=1==1;a{
		t.Log("1==1")
	}
}

func TestSwtichCase(t *testing.T){
	n := 0
	for n < 5{
		switch n {
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("It is not in (0,3)")
		}

		switch {
		case n % 2 == 0:
			t.Log("It is Even")
		case n % 2 == 1:
			t.Log("It is Odd")
		default:
			t.Log("UnKown")
		}
		n++
	}
}
