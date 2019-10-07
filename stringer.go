package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v years", p.Name, p.Age)
}

func main() {
	a := Person{"zhangsan", 34}
	b := Person{"lisi", 54}
	fmt.Println(a)
	fmt.Println(b)
}
