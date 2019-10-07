package main

import "math"
import "fmt"

type Vertex struct {
	X, Y float64
}
type Myfloat float64

type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := Myfloat(math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
}
