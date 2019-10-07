package main

import "fmt"

func main() {
	p := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(p)
	q := []bool{true, false, false, true, true, false}
	fmt.Println(q)

	r := []struct {
		x int
		y bool
	}{{1, true}, {2, false}, {3, false}, {4, false}, {5, true}}

	fmt.Println(r)
}
