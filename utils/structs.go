package utils

type Coord struct {
	X int
	Y int
}

func (c Coord) GetKey() string {
	return ToString(c.X) + "-" + ToString(c.Y)
}
