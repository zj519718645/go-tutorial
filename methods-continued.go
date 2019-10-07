package main

import "fmt"
import "math"

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
