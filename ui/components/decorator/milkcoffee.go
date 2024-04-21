package decorator

type milkCoffee struct {
	coffee iCoffee
}

func (mc *milkCoffee) getCost() float64 {
	coffeePrice := mc.coffee.getCost()
	return coffeePrice + .5
}
