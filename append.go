package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v len=%d cap=%d", s, len(s), cap(s))
}

func main() {
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 3, 4, 5)
	printSlice(s)
}
