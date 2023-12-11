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

var SurroundingCoordsExcDiagonal = [][2]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

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
