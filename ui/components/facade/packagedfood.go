package facade

type PackagedFood struct {
}

func (pf *PackagedFood) newPackagedFood(food *Food) *Food {
	food.contents = food.getFood() + " in a bag"
	return food
}
