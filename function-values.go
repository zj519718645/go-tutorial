package main

import (
	"fmt"
	"math"
)

func computer(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(computer(hypot))
	fmt.Println(computer(math.Pow))
}
