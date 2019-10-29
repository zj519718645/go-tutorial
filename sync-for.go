package main

import "fmt"

var a string
var done bool

func setup() {
	a = "hello world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	fmt.Println(a)
}
