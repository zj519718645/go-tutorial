package main

import "fmt"

func Inc() (v int) {
	defer func() {
		v++
	}()
	return 42
}
func main() {
	fmt.Println(Inc())
}
