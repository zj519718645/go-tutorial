package ch6

import (
	"fmt"
	"math"
	"testing"
)

type shaper interface{}

type circle struct {
	radius float64
}
type rectangles struct {
	width  float64
	length float64
}

func (r rectangles) getDiagonalLength() float64 {
	return math.Sqrt(r.length*r.length + r.width*r.width)
}

func checkAndProcess(s shaper) {
	varR, ok := s.(rectangles)
	if ok {
		fmt.Println(varR.getDiagonalLength())
	} else {
		fmt.Println(varR)
	}

}

func TestTypeAssertion(t *testing.T) {
	var shaperA shaper
	var shaperB shaper
	shaperA = circle{5}
	shaperB = rectangles{3, 5}
	checkAndProcess(shaperA)
	checkAndProcess(shaperB)
}
