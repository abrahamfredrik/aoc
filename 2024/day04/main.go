package main

import (
	"fmt"

	"github.com/abrahamfredrik/aoc/utils"
)

var EXAMPLE = "exampleInput.txt"
var DAY_INPUT = "dayInput.txt"

var letters = []string{"X", "M", "A", "S"}

func main() {
	task1()
	fmt.Println("---")
	task2()
	fmt.Println(utils.SurroundingCoordsIncDiagonal )
}

func task1() {
	input := utils.ReadInput(DAY_INPUT)
	sum := 0
	matrix := utils.CreateMatrix(input)
	for r, row := range matrix {
		for c := range row {
			for _, direction := range utils.SurroundingCoordsIncDiagonal {
				sum += findXmas(matrix, r, c, 0, direction)
			}
		}
	}
	fmt.Println(sum)
}

func task2() {
	input := utils.ReadInput(DAY_INPUT)
	matrix := utils.CreateMatrix(input)
	sum := 0
	for r, row := range matrix {
		for c := range row {
			sum += findMasCross(matrix, r, c)
		}
	}
	fmt.Println(sum)
}

func findMasCross(matrix [][]string, r int, c int) int {
	// Not the prettiest solution but it works :D 
	if !isInside(matrix, r, c) {
		return 0
	}
	if matrix[r][c] != "A" {
		return 0
	}
	diagonals := [][2]int{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	count := 0
	for _, d := range diagonals {
		mCoord := utils.Coord{X: c + d[1], Y: r + d[0]}
		sCoord := utils.Coord{X: c - d[1], Y: r - d[0]}
		if !isInside(matrix, mCoord.Y, mCoord.X) {
			return 0
		}
		if !isInside(matrix, sCoord.Y, sCoord.X) {
			return 0
		}
		if matrix[mCoord.Y][mCoord.X] == "M" && matrix[sCoord.Y][sCoord.X] == "S" {
			count += 1
		}
	}
	if count > 2 {
		panic("How did I even get here?")
	}

	if count == 2 {
		return 1
	} else {
		return 0
	}
}

func isInside(matrix [][]string, r int, c int) bool {
	if r < 0 || (r >= len(matrix)) {
		return false
	}
	if c < 0 || (c >= len(matrix[r])) {
		return false
	}
	return true
}

func findXmas(matrix [][]string, r int, c int, index int, direction [2]int) int {
	if !isInside(matrix, r, c) {
		return 0
	}
	if matrix[r][c] != letters[index] {
		return 0
	}
	if index == len(letters)-1 {
		return 1
	}

	return findXmas(matrix, r+direction[0], c+direction[1], index+1, direction)
}
