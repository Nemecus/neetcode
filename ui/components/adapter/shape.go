package adapter

type shape struct {
}

func (s *shape) startCircleToShape(newShape iShape) {
	newShape.circleToSquareAdapter()
}
