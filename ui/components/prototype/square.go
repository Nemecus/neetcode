package prototype

type square struct {
	width  int
	height int
}

func (s *square) getWidth() int {
	return s.width
}

func (s *square) getHeight() int {
	return s.height
}

func (s *square) clone() shape {
	return &square{width: s.width, height: s.height}
}
