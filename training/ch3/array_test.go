package ch3

import "testing"

func TestArray(t *testing.T) {

	//initialization
	arr1 := [5]int{10, 20, 30, 40, 50}
	arr2 := [...]int{10, 20, 30, 40, 50}
	arr3 := [5]int{1: 10, 2: 20}
	t.Log(arr1, arr2, arr3)

	//access
	var a1 [2]int
	a1[0] = 10
	a1[1] = 20
	for i := 0; i < len(a1); i++ {
		t.Log(i)
	}
	for _, value := range a1 {
		t.Log(value)
	}

}
