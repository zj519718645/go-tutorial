package ch4

import "testing"



func TestMethods(t *testing.T) {
	personA := Person{"John", "Doe", "Male"}
	t.Log(personA.GetName())
	personA.SetForename("Rachard")
	t.Log(personA.GetName())
	personA.SetLastName("Roe")
	t.Log(personA.GetName())
}
