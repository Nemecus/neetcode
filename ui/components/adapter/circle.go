package adapter

type circle struct {
	radius int
}

func (c *circle) getRadius() int {
	return c.radius
}
