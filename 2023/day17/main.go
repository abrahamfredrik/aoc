package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	// day2()
}

func day1() {
	heatLossMatrix, bestValueMatrix := parseInput()
	maxX := len(heatLossMatrix[0]) - 1
	maxY := len(heatLossMatrix) - 1
	goal := utils.Coord{X: len(heatLossMatrix[0]) - 1, Y: len(heatLossMatrix) - 1}
	bestValueMatrix[0][0] = heatLossMatrix[0][0]
	paths := []Path{{[]utils.Coord{utils.Coord{X: 0, Y: 0}}, heatLossMatrix[0][0]}}
	finished := []Path{}
	for len(paths) > 0 {
		current := paths[len(paths)-1]

		paths = paths[:len(paths)-1]
		for _, dxdy := range utils.SurroundingCoordsExcDiagonal {
			if hasHistory(current, dxdy) {
				continue
			}
			c := utils.Coord{X: current.getCurrent().X + dxdy[0], Y: current.getCurrent().Y + dxdy[1]}

			if c.X < 0 || c.X > maxX || c.Y < 0 || c.Y > maxY {
				continue
			}
			if bestValueMatrix[c.Y][c.X] < current.value+heatLossMatrix[c.Y][c.X] {
				// continue
			}
			if slices.Contains(current.history,c){
				continue
			}
			new := Path{[]utils.Coord{}, current.value + heatLossMatrix[c.Y][c.X]}

			new.history = append(new.history, current.history...)

			new.history = append(new.history, c)
			bestValueMatrix[c.Y][c.X] = new.value
			if new.getCurrent() != goal {
				paths = append(paths, new)
			} else {
				temp := []Path{new}
				for _, v := range finished {
					if v.value <= new.value {
						temp = append(temp, v)
					}
				}
				finished = temp
			}
		}
	}
	temp := make([][]string, len(bestValueMatrix))
	for y, line := range bestValueMatrix {
		temp[y] = make([]string, len(line))
		for x := range line {
			temp[y][x] = "."
		}
	}

	for _, v := range finished {
		fmt.Println(v)
		printHistory(v.history, heatLossMatrix)
	}
	// utils.PrintMatrix(temp)
	fmt.Println(bestValueMatrix[maxY][maxX])
}

func printHistory(history []utils.Coord, heatLossMap [][]int) {
	m := map[utils.Coord]string{}
	sum := 0
	for _, v := range history {
		sum += heatLossMap[v.Y][v.X]
		m[v] = "#" // utils.ToString(sum)
	}
	utils.PrintMap(m)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func hasHistory(p Path, dxdy [2]int) bool {
	if len(p.history) < 3 {
		return false
	}
	current := p.getCurrent()
	for i := len(p.history) - 2; i > len(p.history)-4; i-- {
		if !(current.X-dxdy[0] == p.history[i].X && current.Y-dxdy[1] == p.history[i].Y) {
			return false
		}
		current = p.history[i]
	}
	return true
}

func parseInput() ([][]int, [][]int) {
	input := utils.ReadInput("example.txt")
	heatLossMatrix := make([][]int, len(input))
	bestValueMatrix := make([][]int, len(input))
	for y, line := range input {
		heatLossMatrix[y] = make([]int, len(line))
		bestValueMatrix[y] = make([]int, len(line))
		for x, c := range line {
			heatLossMatrix[y][x] = utils.ToInt(c)
			bestValueMatrix[y][x] = math.MaxInt
		}
	}
	return heatLossMatrix, bestValueMatrix
}

type Path struct {
	history []utils.Coord
	value   int
}

func (p *Path) getCurrent() utils.Coord {
	return p.history[len(p.history)-1]
}

type bestValueHistory struct {
	history [3]utils.Coord
	value   int
}
