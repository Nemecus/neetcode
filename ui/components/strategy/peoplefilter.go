package strategy

type PeopleFilter interface {
	apply(person *Person) bool
}
