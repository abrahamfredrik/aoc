package main

import (
	"fmt"

	"github.com/abrahamfredrik/aoc/utils"
)

var maxX int
var maxY int
var m map[utils.Coord]string

func main() {
	parseInput()
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	tiltNorth()
	calculateLoad(m)
}

func day2() {
	// A loop is entered with a cycle of 34 spins

	load := 0
	for i := 0; i < 1000000000%34+(34*5); i++ { // Do 24 spins to get to the starting point and the 5 loops to enter the looping
		spin()
		load = calculateLoad(m)
	}
	fmt.Println(load)
}

func calculateLoad(m map[utils.Coord]string) int {
	load := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			point := utils.Coord{X: x, Y: y}
			if val, ok := (m)[point]; ok {
				if val == "O" {
					load += maxY + 1 - point.Y
				}
			}
		}
	}
	return load
}

func spin() {
	tiltNorth()
	tiltWest()
	tiltSouth()
	tiltEast()
}

func tiltNorth() {
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			point := utils.Coord{X: x, Y: y}
			if val, ok := m[point]; ok {
				if val == "O" {
					delete(m, point)
					above := point.GetAbove()
					temp := m[above]
					for above.Y >= 0 && temp == "" {
						point = above
						above = point.GetAbove()
						temp = m[above]
					}
					m[point] = "O"
				}
			}
		}
	}
}

func tiltSouth() {
	for x := 0; x <= maxX; x++ {
		for y := maxY; y >= 0; y-- {
			point := utils.Coord{X: x, Y: y}
			if val, ok := m[point]; ok {
				if val == "O" {
					delete(m, point)
					below := point.GetBelow()
					for below.Y <= maxY && m[below] == "" {
						point = below
						below = point.GetBelow()
					}
					m[point] = "O"
				}
			}
		}
	}
}

func tiltWest() {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			point := utils.Coord{X: x, Y: y}
			if val, ok := m[point]; ok {
				if val == "O" {
					delete(m, point)
					left := point.GetLeft()
					for left.X >= 0 && m[left] == "" {
						point = left
						left = point.GetLeft()
					}
					m[point] = "O"
				}
			}
		}
	}
}

func tiltEast() {
	for y := maxY; y >= 0; y-- {
		for x := maxX; x >= 0; x-- {
			point := utils.Coord{X: x, Y: y}
			if val, ok := m[point]; ok {
				if val == "O" {
					delete(m, point)
					right := point.GetRight()
					for right.X <= maxX && m[right] == "" {
						point = right
						right = point.GetRight()
					}
					m[point] = "O"
				}
			}
		}
	}
}

func parseInput() {
	input := utils.ReadInput("dayInput.txt")
	m = map[utils.Coord]string{}
	for rowIndex, row := range input {
		for colIndex, col := range row {
			if col == '.' {
				continue
			}
			m[utils.Coord{X: colIndex, Y: rowIndex}] = utils.ToString(col)
		}
	}

	maxX = -1
	maxY = -1
	for k := range m {
		if k.Y > maxY {
			maxY = k.Y
		}
		if k.X > maxX {
			maxX = k.X
		}
	}

}
