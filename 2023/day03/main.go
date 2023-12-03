package main

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	input, matrix := parseInput()
	output := removeInvalidNumbers(input, matrix)

	sum := 0
	for _, numbers := range output {
		numberString := ""
		for _, number := range numbers {
			numberString = numberString + matrix[number.Y][number.X]
		}
		sum += utils.ToInt(numberString)
	}
	fmt.Println(sum)
}

func day2() {
	input, matrix := parseInput()
	output := removeInvalidNumbers2(input, matrix)
	sum := 0
	for _, numbers := range output {
		if len(numbers) == 2 {
			powerRatioNumbers := []int{}
			for _, number := range numbers {
				numberString := ""
				for _, n := range number {
					numberString = numberString + matrix[n.Y][n.X]
				}
				powerRatioNumbers = append(powerRatioNumbers, utils.ToInt(numberString))
			}
			sum += powerRatioNumbers[0] * powerRatioNumbers[1]
		}
	}
	fmt.Println(sum)
}

func removeInvalidNumbers2(input [][]utils.Coord, matrix [][]string) map[utils.Coord][][]utils.Coord {
	gearRegex, _ := regexp.Compile("\\*")
	m := make(map[utils.Coord][][]utils.Coord)
outer:
	for _, numbers := range input {
		for _, number := range numbers {
			for _, dxdy := range utils.SurroundingCoordsIncDiagonal {
				x := number.X + dxdy[0]
				y := number.Y + dxdy[1]
				if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[y]) {
					continue
				}
				if gearRegex.MatchString(matrix[y][x]) {
					if _, ok := m[utils.Coord{X: x, Y: y}]; !ok {
						m[utils.Coord{X: x, Y: y}] = [][]utils.Coord{}
					}
					m[utils.Coord{X: x, Y: y}] = append(m[utils.Coord{X: x, Y: y}], numbers)
					continue outer
				}
			}
		}
	}
	return m
}

func removeInvalidNumbers(input [][]utils.Coord, matrix [][]string) (output [][]utils.Coord) {
	numberAndDotRegex, _ := regexp.Compile("\\.|[0-9]")
outer:
	for _, numbers := range input {
		for _, number := range numbers {
			for _, dxdy := range utils.SurroundingCoordsIncDiagonal {
				x := number.X + dxdy[0]
				y := number.Y + dxdy[1]
				if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[y]) {
					continue
				}
				if !numberAndDotRegex.MatchString(matrix[y][x]) {
					output = append(output, numbers)
					continue outer
				}
			}
		}
	}
	return output
}

func parseInput() ([][]utils.Coord, [][]string) {
	input := utils.ReadInput("dayInput.txt")
	matrix := make([][]string, len(input))
	list := [][]utils.Coord{}
	numberIndex := -1
	for rowIndex, row := range input {
		matrix[rowIndex] = make([]string, len(row))
		partOfNumberOngoing := false
		for colIndex, letter := range row {
			matrix[rowIndex][colIndex] = utils.ToString(letter)
			if unicode.IsDigit(letter) {
				if !partOfNumberOngoing {
					numberIndex++
					list = append(list, []utils.Coord{})
					partOfNumberOngoing = true
				}
				list[numberIndex] = append(list[numberIndex], utils.Coord{X: colIndex, Y: rowIndex})
			} else {
				partOfNumberOngoing = false
			}
		}
	}
	return list, matrix
}
