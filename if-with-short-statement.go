package main

import "fmt"
import "math"

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 3, 25), pow(2, 4, 18))
}
