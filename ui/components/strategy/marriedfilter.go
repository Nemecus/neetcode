package strategy

type MarriedFilter struct {
}

func (mf *MarriedFilter) apply(person *Person) bool {
	return person.isMarried()
}
