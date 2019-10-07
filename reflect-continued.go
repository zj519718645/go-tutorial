package main

import (
	"fmt"
	"reflect"
)

type cat struct{}
type Enum int

const (
	Zero Enum = 0
)

func main() {
	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	typeOfA := reflect.TypeOf(Zero)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
