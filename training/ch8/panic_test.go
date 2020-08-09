package ch8

import (
	"errors"
	"fmt"
	"testing"
)

func DivideTest02(dividend float64, divisor float64) (result float64, err error) {
	if divisor == 0 {
		result = 0.0
		err = errors.New("runtime error: divide by zero")
		panic(err)
		return
	}
	result = divisor / divisor
	err = nil
	return
}

func TestPanic(t *testing.T) {
	r, err := DivideTest02(6.6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生panic异常！错误信息为：", err)
		}
	}()
}
