//bool
//string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//float32 float64
//complex64 complex128
package main

import "fmt"

var (
	a int   = 11
	b int8  = 128
	c int16 = 115
	d int32 = 4455
	e int64 = 111
)

func main() {
	fmt.Println(a, b, c, d, e)
}
