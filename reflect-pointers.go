package main

import "fmt"
import "reflect"

type cat struct{}

func main() {
	v := &cat{}
	typeOfCat := reflect.TypeOf(v)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	//fmt.Println(typeOfCat.Name())
	typeOfCat = typeOfCat.Elem()
	fmt.Println(typeOfCat.Name(),typeOfCat.Kind())
}
