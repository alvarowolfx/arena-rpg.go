package lib

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) *Position {
	return &Position{X: x, Y: y}
}
