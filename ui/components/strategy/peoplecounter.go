package strategy

type PeopleCounter struct {
	filter PeopleFilter
}

func (pc *PeopleCounter) setFilter(filter PeopleFilter) {
	pc.filter = filter
}

func (pc *PeopleCounter) count(people People) int {
	count := 0
	for i := 0; i < len(people.people); i++ {
		if pc.filter.apply(&people.people[i]) {
			count++
		}
	}
	return count
}
