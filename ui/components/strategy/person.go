package strategy

type Person struct {
	lastName string
	age      int
	married  bool
}

func (p *Person) getLastName() string {
	return p.lastName
}

func (p *Person) getAge() int {
	return p.age
}

func (p *Person) isMarried() bool {
	return p.married
}
