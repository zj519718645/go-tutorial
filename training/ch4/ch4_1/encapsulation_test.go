package ch4_1

import (
	"testing"

	"../../ch4"
)

func TestEncapsulation1(t *testing.T) {
	personB := ch4.Person{"Jennyfer", "Doe", "Female"}
	t.Log(personB.forename)
	t.Log(personB.Sex)
}

func TestEncapsulation2(t *testing.T) {
	//var personB = new(ch4.Person)
	var personB ch4.Person
	t.Log(personB)
	personB.Sex = "Female"
	personB.SetForename("li")
	personB.SetLastName("si")
	t.Log(personB)
}
