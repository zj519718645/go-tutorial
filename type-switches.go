package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}
}

func main() {
	do(21)
	do("hello")
	do(234.34)
}
