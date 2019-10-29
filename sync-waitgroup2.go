package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ws *sync.WaitGroup, ch chan int) {
	defer ws.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ch:
			return
		}
	}
}

func main() {
	var ws sync.WaitGroup
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		ws.Add(1)
		go worker(&ws, ch)
	}
	time.Sleep(time.Second)
	close(ch)
	ws.Wait()
}
