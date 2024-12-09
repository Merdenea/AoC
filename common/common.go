package common

type Pair struct {
	I, J int64
}

type Direction Pair

var (
	UP    = Direction{-1, 0}
	DOWN  = Direction{1, 0}
	LEFT  = Direction{0, -1}
	RIGHT = Direction{0, 1}
)

func TurnRight(d Direction) Direction {
	switch d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	}
	return UP
}
