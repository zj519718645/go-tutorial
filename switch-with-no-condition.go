package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoo")
	default:
		fmt.Println("good evening")
	}
}
