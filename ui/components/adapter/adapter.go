package adapter

type circleAdapter struct {
	circleItem *circle
}

func (ca *circleAdapter) circleToSquareAdapter() square {
	newSquare := (*square).createSquare(&square{}, (ca.circleItem.radius * 2))

	return newSquare
}
