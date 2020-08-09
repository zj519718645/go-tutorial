package ch3

import "testing"

func TestPointer(t *testing.T) {
	var ptrA *int
	var ptrB = new(float32)
	var c float64

	ptrC := &c
	t.Log(ptrA, ptrB, ptrC)
	t.Log(*ptrB)

	var a = 10
	var p = &a
	t.Log(*p)
	*p = 666
	t.Log(a)
}

func TestNew(t *testing.T) {
	var p1 *int
	p1 = new(int)
	t.Log(*p1)

	p2 := new(int)
	*p2 = 111
	t.Log(*p2)
}

func TestArrayPointer(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	var p *[5]int = &a
	for index, value := range *p {
		t.Log(index, value)
	}
}

func TestPointerArray(t *testing.T) {
	var p2 [5]*int
	i, j := 10, 20
	p2[0] = &i
	p2[1] = &j
	for index, value := range p2 {
		if value != nil {
			t.Log(index, *value)
		}
	}
}
