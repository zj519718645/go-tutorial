package main

import "fmt"

func main() {
	var s []int
	var s1 [10]int
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))

	if s == nil {
		fmt.Println("nil!")
	}
}
