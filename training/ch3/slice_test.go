package ch3

import "testing"

func TestSlice(t *testing.T) {
	//base on array
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	t.Log(mySlice)

	mySlice2 := make([]int, 5)
	for _, data := range mySlice2 {
		t.Log(data)
	}

	mySlice3 := []int{10, 20, 30, 40, 50}
	newSlice := mySlice3[1:3]
	newSlice[1] = 25
	t.Log(len(newSlice), cap(newSlice))
	t.Log(newSlice)
	t.Log(mySlice3)

	newSlice = append(newSlice, 60)
	t.Log(len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 70)
	t.Log(len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 80)
	t.Log(len(newSlice), cap(newSlice))
}
