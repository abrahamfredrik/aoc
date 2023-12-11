package main

import (
	"fmt"
	"math"

	"github.com/abrahamfredrik/aoc/utils"
	"golang.org/x/exp/slices"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	stars := parseInput()
	stars = extendStars(stars, 2)
	dist := findDistence(stars)
	fmt.Println(dist)
}

func day2() {
	stars := parseInput()
	stars = extendStars(stars, 1000000)
	dist := findDistence(stars)
	fmt.Println(dist)
}

func findDistence(stars []utils.Coord) int {
	dist := 0
	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			dist += int(math.Abs(float64(stars[i].X - stars[j].X)))
			dist += int(math.Abs(float64(stars[i].Y - stars[j].Y)))
		}
	}
	return dist
}

func extendStars(inputStars []utils.Coord, extension int) []utils.Coord {
	slices.SortFunc(inputStars, sortX)
	tempStars := []utils.Coord{}
	maxX := inputStars[len(inputStars)-1].X
	ext := 0
	for i := 0; i <= maxX; i++ {
		foundStar := false
		for _, s := range inputStars {
			if i == s.X {
				tempStars = append(tempStars, utils.Coord{X: s.X + ext, Y: s.Y})
				foundStar = true
			}
		}
		if !foundStar {
			ext += extension - 1
		}

	}
	slices.SortFunc(tempStars, sortY)
	outputStars := []utils.Coord{}
	maxY := tempStars[len(tempStars)-1].Y
	ext = 0
	for i := 0; i <= maxY; i++ {
		foundStar := false
		for _, s := range tempStars {
			if i == s.Y {
				outputStars = append(outputStars, utils.Coord{X: s.X, Y: s.Y + ext})
				foundStar = true
			}
		}
		if !foundStar {
			ext += extension - 1
		}
	}

	return outputStars
}

func sortX(a utils.Coord, b utils.Coord) int {
	return a.X - b.X
}

func sortY(a utils.Coord, b utils.Coord) int {
	return a.Y - b.Y
}

func parseInput() []utils.Coord {
	input := utils.ReadInput("dayInput.txt")
	starList := []utils.Coord{}
	for indexRow, row := range input {
		for indexCol, col := range row {
			if col != '.' {
				starList = append(starList, utils.Coord{X: indexCol, Y: indexRow})
			}
		}
	}
	return starList
}
