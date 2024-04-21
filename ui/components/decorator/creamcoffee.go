package decorator

type creamCoffee struct {
	coffee iCoffee
}

func (cc *creamCoffee) getCost() float64 {
	coffeePrice := cc.coffee.getCost()
	return coffeePrice + .7
}
