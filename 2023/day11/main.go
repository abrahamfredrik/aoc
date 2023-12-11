package main

import (
	"fmt"
	"math"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	// day2()
}

func day1() {
	matrix := parseInput()
	matrix = extendMatrix(matrix)
	//utils.PrintMatrix(matrix)
	stars := []utils.Coord{}
	for indexRow, row := range matrix {
		for indexCol, col := range row {
			if col != '.' {
				stars = append(stars, utils.Coord{indexCol, indexRow})
			}
		}
	}

	dist := 0
	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			dist += int(math.Abs(float64(stars[i].X - stars[j].X)))
			dist += int(math.Abs(float64(stars[i].Y - stars[j].Y)))
		}
	}

	fmt.Println(dist)


}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func extendMatrix(inputMatrix [][]rune) [][]rune {
	hasStarRow := make([]bool, len(inputMatrix))
	hasStarCol := make([]bool, len(inputMatrix[0]))
	for indexRow, row := range inputMatrix {
		for indexCol, col := range row {
			if col != '.' {
				hasStarCol[indexCol] = true
				hasStarRow[indexRow] = true
			}
		}
	}
	counterRow := 0
	for _, v := range hasStarRow {
		if !v {
			counterRow++
		}
	}
	outputMatrix := make([][]rune, 0)
	for rowIndex, row := range inputMatrix {
		tempRow := []rune{}
		for colIndex, col := range row {
			tempRow = append(tempRow, col)
			if !hasStarCol[colIndex] {
				tempRow = append(tempRow, '.')
			}
		}
		outputMatrix = append(outputMatrix, tempRow)
		if !hasStarRow[rowIndex] {
			tempRow := []rune{}
			for i := 0; i <= len(inputMatrix[rowIndex])+counterRow; i++ {
				tempRow = append(tempRow, '.')
			}
			outputMatrix = append(outputMatrix, tempRow)
		}
	}

	return outputMatrix
}

func parseInput() [][]rune {
	input := utils.ReadInput("dayInput.txt")
	matrix := make([][]rune, len(input))
	for indexRow, row := range input {
		matrix[indexRow] = make([]rune, len(row))
		for indexCol, col := range row {
			matrix[indexRow][indexCol] = col
		}
	}
	return matrix
}
