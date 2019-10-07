package main

import "fmt"

type I interface{}

func main() {
	var i I = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64)
	//fmt.Println(f)
}
