package main

import "fmt"
import "strconv"

func main() {
	var orig = "666"
	var an int = 299
	var f float64 = 34.55
	var newS string = "67.34"
	//fmt.Println(strconv.IntSize)
	fmt.Println(strconv.Itoa(an))
	fmt.Println(strconv.FormatFloat(f, 'f', 2, 64))
	fmt.Println(strconv.Atoi(orig))
	fmt.Println(strconv.ParseFloat(newS, 64))

}
