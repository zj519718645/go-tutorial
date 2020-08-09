package ch4

import (
	"fmt"
	"testing"
)

func Myfuc() {
	a := 1
	fmt.Println("a=", a)
}

func test(a int) {
	b := 2
	s := a + b
	fmt.Println(s)
}

func Myfuc2() (result int) {
	result = 2
	return
}

func Myfuc3() (i float32, j int, k int) {
	i, j, k = 11, 12, 14
	return
}

func test2(a ...int) {
	fmt.Println(len(a))
	for _, value := range a {
		fmt.Println(value)
	}
}

func TestFunctions(t *testing.T) {
	Myfuc()
	test(1)
	t.Log(Myfuc2())
	v1, v2, v3 := Myfuc3()
	t.Log(v1, v2, v3)
	test2(2, 3, 4)
}
