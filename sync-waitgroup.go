package main

import (
	"fmt"
	"sync"
)

func main() {
	ws := sync.WaitGroup{}
	ws.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			ws.Done()
		}(i)
		
	}
	ws.Wait()
	fmt.Println("hello world")
}
