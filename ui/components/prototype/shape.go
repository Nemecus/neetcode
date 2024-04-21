package prototype

type shape interface {
	getWidth() int
	getHeight() int
	clone() shape
}
