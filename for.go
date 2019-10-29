package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	i := 0
	for i < 10 {
		fmt.Println(i)
	}

	for j := 0; j< 10; j++ {
		fmt.Println(j)
	}
}
