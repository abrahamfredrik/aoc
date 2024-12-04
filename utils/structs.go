package utils

var SurroundingCoordsExcDiagonal = [][2]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

var Diagonals = [][2]int{
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

var SurroundingCoordsIncDiagonal = append(SurroundingCoordsExcDiagonal, Diagonals...)

type Coord struct {
	X int
	Y int
}

func (c Coord) GetKey() string {
	return ToString(c.X) + "-" + ToString(c.Y)
}

func (c Coord) GetSurroundingExcDiagnonal() []Coord {
	return c.getSurrounding(SurroundingCoordsExcDiagonal)
}

func (c Coord) GetSurroundingIncDiagnonal() []Coord {
	return c.getSurrounding(SurroundingCoordsIncDiagonal)
}

func (c Coord) getSurrounding(input [][2]int) []Coord {
	coords := []Coord{}
	for _, diff := range input {
		coords = append(coords, Coord{c.X + diff[0], c.Y + diff[1]})
	}
	return coords
}

func (c Coord) GetAbove() Coord {
	return Coord{c.X, c.Y - 1}
}

func (c Coord) GetBelow() Coord {
	return Coord{c.X, c.Y + 1}
}

func (c Coord) GetRight() Coord {
	return Coord{c.X+1, c.Y}
}

func (c Coord) GetLeft() Coord {
	return Coord{c.X-1, c.Y}
}
