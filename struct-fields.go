package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vt := Vertex{1, 2}
	vt.X = 3
	fmt.Println(vt.X, vt.Y)
}
