package main

import "fmt"
import "time"

func worker(ch chan int) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ch:
		}
	}
}

func main() {
	ch := make(chan int)
	for i:=0; i<10; i++ {
		go worker(ch)
	}
	time.Sleep(time.Second)
	close(ch)
}