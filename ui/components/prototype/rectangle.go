package prototype

type rectangle struct {
	width  int
	height int
}

func (r *rectangle) getWidth() int {
	return r.width
}

func (r *rectangle) getHeight() int {
	return r.height
}

func (r *rectangle) clone() shape {
	return &rectangle{width: r.width, height: r.height}
}
