package facade

type Chef struct {
}

func (c *Chef) prepareFood(newOrder *Order) *Food {
	return &Food{
		contents: newOrder.contents,
	}
}
