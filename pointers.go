package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)

	*p = 45
	fmt.Println(i)

	p = &j
	fmt.Println(*p)

	*p = 444
	fmt.Println(j)
}
