package decorator

type sugarCoffee struct {
	coffee iCoffee
}

func (sc *sugarCoffee) getCost() float64 {
	coffeePrice := sc.coffee.getCost()
	return coffeePrice + .2
}
