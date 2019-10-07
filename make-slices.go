package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)
}

func printSlice(s []int) {
	fmt.Printf("%v len is %d cap is %d", s, len(s), cap(s))
}
