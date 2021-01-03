package main

import "flag"

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

}
