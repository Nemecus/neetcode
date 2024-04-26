package strategy

type SeniorFilter struct {
}

func (sf *SeniorFilter) apply(person *Person) bool {
	return person.getAge() >= 65
}
