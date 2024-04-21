package builder

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildMeal(cost float64, takeout bool, main string, drink string) meal {
	d.builder.addCost(cost)
	d.builder.addTakeout(takeout)
	d.builder.addMainCourse(main)
	d.builder.addDrink(drink)
	return d.builder.build()
}
