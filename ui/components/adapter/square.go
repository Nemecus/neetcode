package adapter

type square struct {
	sideLength int
}

func (s *square) getSideLength() int {
	return s.sideLength
}
