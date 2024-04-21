package builder

type iBuilder interface {
	addCost(float64)
	addTakeout(bool)
	addMainCourse(string)
	addDrink(string)
	build() meal
}

func getBuilder() iBuilder {
	return newMeal()
}
