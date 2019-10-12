package main

import "sync"
import "sync/atomic"

var total uint64

func worker(ws *sync.WaitGroup) {
	defer ws.Done()
	for i := 0; i < 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
}
