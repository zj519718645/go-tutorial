package ch3

import "testing"

func TestStruct(t *testing.T) {
	type Student struct {
		ID      int
		Name    string
		Address string
		Age     int
		Class   string
		Teacher string
	}
	var lm Student
	lm.ID = 20200411
	lm.Class = "物联网工程1801"
	lm.Age = 18
	lm.Address = "山西省太原市"
	lm.Teacher = "蕾蕾"
	lm.Name = "黎明"
	t.Log(lm)

	var s Student
	var p1 *Student
	p1 = &s
	t.Log(s)
	t.Log(*p1)

	p2 := new(Student)
	p2.ID = 202020311
	p2.Name = "李四"
	t.Log(p2)
	t.Log(p2.Name)
}

func TestInnerStruct(t *testing.T) {
	type Pointer struct {
		X, Y int
	}

	type Circle struct {
		Pointer
		Radius int
	}

	type Cylinder struct {
		Pointer
		Height int
	}

	var c Cylinder
	c.X = 10
	c.Y = 10
	c.Radius = 10
	c.Height = 10

}
