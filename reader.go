package main

import "fmt"
import "io"
import "strings"

func main() {
	rd := strings.NewReader("Hello, World!")
	b := make([]byte, 8)
	for {
		n, err := rd.Read(b)
		fmt.Printf("n = %v err = %v b = %v", n, err, b)
		if err == io.EOF {
			break
		}
	}
}
