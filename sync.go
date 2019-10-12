package main

import "fmt"

func main() {
	c := make(chan bool, 1)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}
	for i := 0; i < 100; i++ {
		<-c
	}
}
