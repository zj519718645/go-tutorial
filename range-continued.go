package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i, _ := range pow {
		fmt.Println(i)
	}
	for _, val := range pow {
		fmt.Println(val)
	}
}
