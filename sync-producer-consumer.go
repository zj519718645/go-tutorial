package main

import (
	"fmt"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	Producer(3, ch)
	Producer(5, ch)
	Consumer(ch)
	time.Sleep(5 * time.Second)
}
