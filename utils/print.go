package utils

import (
	"fmt"
	"math"
)

func PrintMatrix[T string | int | rune](m [][]T) {
	for _, row := range m {
		for _, col := range row {
			l := ToString(col)
			fmt.Print(l)
		}
		fmt.Println()
	}
}

func PrintMap(m map[Coord]string) {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for k := range m {
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
		if k.X < minX {
			minX = k.X
		}
		if k.Y < minY {
			minY = k.Y
		}
	}
	if minY < 0 || minX < 0 {
		panic("Negative numbers not supported (currently)")
	}
	arrayMatrix := make([][]string, maxY+1)
	for i := range arrayMatrix {
		arrayMatrix[i] = make([]string, maxX+1)
	}
	for k, v := range m {
		arrayMatrix[k.Y][k.X] = v
	}
	PrintMatrix(arrayMatrix)
}
