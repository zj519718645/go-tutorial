package main

import (
	"log"

	"github.com/zj519718645/go-tutorial/projects/sqltostruct/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v", err)
	}
}
