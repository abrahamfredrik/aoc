package utils

var SurroundingCoordsIncDiagonal = [][2]int{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

type Coord struct {
	X int
	Y int
}

func (c Coord) GetKey() string {
	return ToString(c.X) + "-" + ToString(c.Y)
}
