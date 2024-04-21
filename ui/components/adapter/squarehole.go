package adapter

type squareHole struct {
	sideLenght int
}

func (sh *squareHole) canFit(box square) bool {
	return sh.sideLenght >= box.getSideLength()
}
