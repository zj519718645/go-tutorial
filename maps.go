//映射的零值为 nil 。nil 映射既没有键，也不能添加键。
//make 函数会返回给定类型的映射，并将其初始化备用。
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	if m == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("not nil")
	}
	m["bell"] = Vertex{19.5, 11.5}
	fmt.Println(m["bell"])
}
