package ch6

import (
	"fmt"
	"testing"
)

func CheckAndProcess(c interface{}) {
	switch varC := c.(type) {
	case int:
		fmt.Println(varC + 1)
	case string:
		fmt.Println("hello ", c)
	default:
		fmt.Println("unknown type")
	}
}

func TestTypeCheck(t *testing.T) {
	CheckAndProcess(5)
	CheckAndProcess("China")
}
