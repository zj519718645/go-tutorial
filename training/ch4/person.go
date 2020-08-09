package ch4

type Person struct {
	forename string
	lastName string
	Sex      string
}

func (p Person) GetName() string {
	return p.forename + " " + p.lastName
}

func (p *Person) SetForename(newForename string) {
	p.forename = newForename
}

func (p Person) SetLastName(newLastName string) {
	p.lastName = newLastName
}
