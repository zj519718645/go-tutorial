package ch7

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Hello from goroutine", i)
		}(i)
	}
	time.Sleep(time.Millisecond)
}
