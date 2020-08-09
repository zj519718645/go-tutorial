package ch6

import (
	"fmt"
	"sync"
	"testing"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1

	m.Unlock()
	wg.Done()
}

func TestWithoutMutex(t *testing.T) {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
