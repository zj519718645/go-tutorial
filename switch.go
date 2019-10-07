package main

import "fmt"
import "runtime"

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS.X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
