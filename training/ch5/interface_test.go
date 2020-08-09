package ch5

import (
	"math"
	"testing"
)

type shape interface {
	area() float64
}

type line interface {
	length() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) length() float64 {
	return math.Pi * c.radius * 2
}

func TestInterface1(t *testing.T) {
	circleA := circle{radius: 10}
	var shapeA shape = &circleA
	t.Log(shapeA.area())
	var shapeB = new(circle)
	shapeB.radius = 10
	t.Log(shapeB.area())
	var lineA = new(circle)
	lineA.radius = 10
	t.Log(lineA.length())
}

type Any interface{}

type retangle struct {
	width  float64
	length float64
}
type shaper interface {
	area() float64
}

func (r *retangle) area() float64 {
	return r.length * r.width
}
func TestInterface2(t *testing.T) {
	var any Any
	var retangleA retangle = retangle{10, 10}
	var shaperA shaper
	any = retangleA
	t.Log(any)
	any = shaperA
	t.Log(any)
	any = &retangleA
	t.Log(any)
	any = 39
	t.Log(any)
	any = "everything"
	t.Log(any)
}
