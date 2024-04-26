package facade

type Food struct {
	contents string
}

func (f *Food) getFood() string {
	return f.contents
}
