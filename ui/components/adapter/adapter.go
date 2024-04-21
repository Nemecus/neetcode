package adapter

type squareAdapter struct {
	squareItem *square
}

func (sa *squareAdapter) circleToSquareAdapter(cir circle) square {
	length := cir.getRadius()
	newSquare := sa.squareItem{sideLength: length}
	return newSquare
}
