/*
继承
一个结构体嵌到另一个结构体，称作组合
匿名和组合的区别
如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
*/
package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (car *Car) run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	lunzi int
}
type Train struct {
	Car
}

func main() {
	c := Car{200, "cacaca"}
	c.run()

	var bk Bike
	bk.weight = 100
	bk.name = "xiwen"
	bk.lunzi = 2
	bk.run()

	var tr Train
	tr.weight = 2000
	tr.name = "high speed"
	tr.run()
}
