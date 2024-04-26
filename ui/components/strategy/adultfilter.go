package strategy

type AdultFilter struct {
}

func (af *AdultFilter) apply(person *Person) bool {
	return person.getAge() >= 18
}
