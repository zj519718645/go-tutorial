package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["answer"] = 4
	fmt.Println(m["answer"])
	m["answer"] = 42
	fmt.Println(m["answer"])
	delete(m, "answer")
	v, ok := m["answer"]
	fmt.Println("The value:", v, "present?", ok)
}
