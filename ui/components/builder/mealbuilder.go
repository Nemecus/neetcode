package builder

type mealBuilder struct {
	cost    float64
	takeout bool
	main    string
	drink   string
}

func newMeal() *mealBuilder {
	return &mealBuilder{}
}

func (mb *mealBuilder) addCost(amount float64) {
	mb.cost = amount
}

func (mb *mealBuilder) addTakeout(value bool) {
	mb.takeout = value
}

func (mb *mealBuilder) addMainCourse(value string) {
	mb.main = value
}

func (mb *mealBuilder) addDrink(value string) {
	mb.drink = value
}

func (mb *mealBuilder) build() meal {
	return meal{
		cost:    mb.cost,
		takeout: mb.takeout,
		main:    mb.main,
		drink:   mb.drink,
	}
}
