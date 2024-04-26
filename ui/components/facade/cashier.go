package facade

type Cashier struct {
}

func (c *Cashier) takeOrder(contents string, takeOut bool) *Order {
	return &Order{
		contents: contents,
		takeOut:  takeOut,
	}
}
