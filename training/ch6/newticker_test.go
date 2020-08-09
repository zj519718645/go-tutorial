package ch6

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
