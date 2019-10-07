package main

import "fmt"

type T struct {
	S string
}
type I interface {
	M()
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t
	i.M()
}
