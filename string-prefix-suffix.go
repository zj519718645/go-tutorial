package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "this is example of string"
	var prefix = "that"

	var suffix = "string"
	fmt.Println(strings.HasPrefix(s, prefix))
	fmt.Println(strings.HasSuffix(s, suffix))
}
