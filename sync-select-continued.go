package main

import "fmt"
import "time"

func worker(channel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-channel:
			//return
		}
	}
}

func main() {
	ch := make(chan bool)
	go worker(ch)
	time.Sleep(time.Second)
	ch <- true
}
