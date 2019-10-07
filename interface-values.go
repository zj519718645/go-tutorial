package main

import "fmt"
import "math"

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func (t *T) M() {
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("%v,%T\n", i, i)
}

func main() {
	var i I
	i = &T{"hello"}
	i.M()
	describe(i)
	i = F(math.Pi)
	i.M()
	describe(i)
}
