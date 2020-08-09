package ch1

import (
	"testing"
)

// 变量
func TestVariables(t *testing.T) {

	/*------------------------声明---------------------------*/
	var v1 int
	var v2 string
	var v3 [10]int
	var v4 []int
	var v5 struct {
		f int
	}
	var v6 *int
	var v7 map[string]int
	var v8 func(a int) int
	v11 := v1
	v12 := v2
	/*-------------------------初始化----------------------*/
	var v13 = 100
	v14 := 0.67
	/*---------------------匿名变量------------------------*/
	_, v15 := 34, 35

	t.Logf("%T,%T,%T,%T,%T,%T,%T,%T,%T,%T,%T,%T,%T\n", v1, v2, v3, v4, v5, v6, v7, v8, v11, v12, v13, v14, v15)

}

func TestConstants(t *testing.T) {
	//显式定义
	const a string = "abc"
	//隐式定义
	const b = "bcd"
	const c = 10
	t.Log(a, b, c)
	//预定义的常亮true/false/iota
	t.Log(true, false)
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	t.Log(a1, a2, a3)
	const (
		b1 = iota
		b2
		b3
	)
	t.Log(b1, b2, b3)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	t.Log(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func TestDataType(t *testing.T) {
	//布尔型
	var b bool
	t.Log(b)

	//整型
	var value2 int32
	value1 := 64
	t.Log(value1, value2)

	//浮点型
	var fvalue1 float32
	fvalue2 := 12.0
	t.Log(fvalue1, fvalue2)

	//复数型
	var v1 complex64
	v2 := 3.2 + 12i
	v3 := complex(3.2, 12)
	t.Log(v1, v2, v3)

	//字符串类型
	var str1 = "hello world"
	t.Log(str1)

	//派生类型
	//1.pointer
	var pb int
	c := &pb
	t.Log(*c)
	var pa = new(int)
	var d = 1
	pa = &d
	t.Log(*pa)
	//2.array
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	for index, value := range arr1 {
		t.Log(index)
		t.Log(value)
	}
	var arr2 = [3]int{2, 4, 6}
	t.Log(arr2)
	var arr3 = [...]int{1, 3, 7}
	t.Log(arr3)
	//3.slice
	var slice1 = []int{1, 3, 7}
	t.Log(slice1)
	var slice2 = arr2[0:2]
	t.Log(slice2)
	//4.map
	var ma = map[int]int{1: 1, 2: 2, 3: 3}
	ma[4] = 4
	t.Log(ma)
	for key, value := range ma {
		t.Log(key, value)
	}
	//5.channel
	ch := make(chan int, 3)
	ch <- 1
	n := 2
	ch <- n
	t.Log(len(ch), cap(ch))
	a := <-ch
	t.Log(a)
	t.Log(<-ch)
	//6.struct
	type Person struct {
		Name  string
		Age   int
		Email string
	}
	var p Person
	p.Name = "zhangsan"
	p.Age = 28
	t.Log(p)
	lisi := Person{
		Name:  "lisi",
		Age:   23,
		Email: "lisi@qq.com",
	}
	t.Log(lisi)
	//7.interface
	var iv interface{} = 250
	switch iv.(type) {
	case int:
		t.Logf("%d is int\n", iv)
	case string:
		t.Logf("%s is string\n", iv)
	default:
		t.Error("unknown type")
	}
}
