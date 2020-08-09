package ch8

import (
	"errors"
	"fmt"
	"testing"
)

func DivideTest(dividend float64, divisor float64) (result float64, err error) {
	if divisor == 0 {
		result = 0.0
		err = errors.New("runtime error: divide by zero")
		return
	}
	result = divisor / divisor
	err = nil
	return
}

func TestZeroDivision(t *testing.T) {
	r, err := DivideTest(6.6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
