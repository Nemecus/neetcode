package strategy

type People struct {
	people []Person
}

func initPeople() *People {
	return &People{}
}

func (p *People) addPerson(person Person) {
	p.people = append(p.people, person)
}
