package ch2

import "testing"

func TestIfControl(t *testing.T) {
	if a := 3; a == 3 {
		t.Log("a == 3")
	}
	if a := 5; a == 10 {
		t.Log("equals")
	} else if a > 10 {
		t.Log("gt")
	} else {
		t.Log("lt")
	}
}

func TestSwitch(t *testing.T) {
	num := 3
	switch num {
	case 1:
		t.Log("floor 1")
	case 2:
		t.Log("floor 2")
	case 3:
		t.Log("floor 3")
	default:
		t.Log("floor X")
	}

	switch score := 65; {
	case score > 90:
		t.Log("A+")
	case score > 80:
		t.Log("A")
	case score > 70:
		t.Log("B+")
	case score > 60:
		t.Log("B")
	default:
		t.Log("below Peer")
	}
}

func TestLoop(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	t.Log(sum)
	str1 := "abcc"
	for index, value := range str1 {
		t.Log(index, value)
	}
}
