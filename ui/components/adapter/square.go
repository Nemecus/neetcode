package adapter

type square struct {
	sideLength int
}

func (s *square) createSquare(length int) square {
	return square{sideLength: length}
}

func (s *square) getSideLength() int {
	return s.sideLength
}
