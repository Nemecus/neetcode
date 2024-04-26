package facade

type Order struct {
	contents string
	takeOut  bool
}

func newOrder(items string, toGo bool) *Order {
	return &Order{
		contents: items,
		takeOut:  toGo,
	}
}

func (o *Order) getOrder() string {
	return o.contents
}

func (o *Order) isTakeOut() bool {
	return o.takeOut
}
