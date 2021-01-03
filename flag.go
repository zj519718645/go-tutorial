package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go语言编程", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}
