package ch6

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func onceFunc() {
	fmt.Println("Only run once")

}

func TestOnce(t *testing.T) {
	var once sync.Once
	for i := 0; i < 10; i++ {
		j := i
		go func(int) {
			once.Do(onceFunc)
			fmt.Println(j)
		}(j)
	}
	time.Sleep(time.Second * 2)
}
