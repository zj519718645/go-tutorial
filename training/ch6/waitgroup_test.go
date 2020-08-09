package ch6

import (
	"fmt"
	"sync"
	"testing"
)

var wt sync.WaitGroup

func goFunc(i int) {
	fmt.Println(i)
	wt.Done()
}

func TestWaitGroup(t *testing.T) {
	for i := 0; i < 10; i++ {
		wt.Add(1)
		go goFunc(i)
	}
	wt.Wait()
	fmt.Println("此处代码在协程执行完毕后输出")
}
