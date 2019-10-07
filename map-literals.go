package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"bell": Vertex{
			111.33, 222.434,
		},
		"google": Vertex{
			55.3, 5455.5,
		},
	}
	fmt.Println(m)
}
