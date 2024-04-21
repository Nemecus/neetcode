package decorator

type simpleCoffee struct{}

func (sc *simpleCoffee) getCost() float64 {
	return 1.1
}
