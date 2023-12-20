package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

var m map[utils.Coord]string

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	current := parseInput()
	visited := map[utils.Coord]bool{current: true}
	next := findFirstSteps(current)[0] // Select one direction

	for m[next] != "S" {
		paths := getPaths(next)
		if paths[0] == current {
			current = next
			next = paths[1]
		} else {
			current = next
			next = paths[0]
		}
		visited[current] = true
	}

	fmt.Println(len(visited) / 2)
}

func day2() {
	current := parseInput()
	visited := map[utils.Coord]bool{current: true}
	next := findFirstSteps(current)[0] // Select one direction

	for m[next] != "S" {
		paths := getPaths(next)
		if paths[0] == current {
			current = next
			next = paths[1]
		} else {
			current = next
			next = paths[0]
		}
		visited[current] = true
	}

	fmt.Println("1--")
	utils.PrintMap(m)
	s := parseInput()
	m[s] = "7"
	fmt.Println("2--")
	utils.PrintMap(m)
	fmt.Println("3--")
	for k := range m {
		if visited[k] {
			continue
		} else if m[k] != "O" && m[k] != "I" {
			m[k] = "."
		}
		counter := 0
		enterWall := ""
		for i := k.Y - 1; i >= s.Y; i-- {
			if strings.ContainsAny(m[utils.Coord{X: k.X, Y: i}], "-") {
				counter++
			}
			if strings.ContainsAny(m[utils.Coord{X: k.X, Y: i}], "JL") {
				enterWall = m[utils.Coord{X: k.X, Y: i}]
			}
			if strings.ContainsAny(m[utils.Coord{X: k.X, Y: i}], "F") {
				if enterWall == "J" {
					counter++
				}
			}
			if strings.ContainsAny(m[utils.Coord{X: k.X, Y: i}], "7") {
				if enterWall == "L" {
					counter++
				}
			}
		}
		if counter%2 == 1 {
			m[k] = "I"
		} else {
			m[k] = "O"
		}
	}

	counter := 0
	for _, v := range m {
		if v == "I" {
			counter++
		}
	}
	utils.PrintMap(m)
	fmt.Println(counter)
}

func bubbleUp(c utils.Coord) bool {
	if m[c] == "" {
		return true
	}
	if m[c] != "." {
		return false
	}
	m[c] = "*"
	around := c.GetSurroundingExcDiagnonal()
	outside := false
	for _, v := range around {
		outside = bubbleUp(v) || outside
	}
	return outside
}

func findFirstSteps(s utils.Coord) []utils.Coord {
	possiblePaths := []utils.Coord{}
	for _, c := range s.GetSurroundingExcDiagnonal() {
		if m[c] == "." || m[c] == "" {
			continue
		}
		endings := getPaths(c)
		for _, v := range endings {
			if v == s {
				possiblePaths = append(possiblePaths, c)
			}
		}
	}
	if len(possiblePaths) != 2 {
		panic("More than 2 start point")
	}
	return possiblePaths
}

func getPaths(coord utils.Coord) []utils.Coord {
	s := m[coord]
	diff := []utils.Coord{}
	if s == "-" {
		diff = []utils.Coord{
			{1, 0},
			{-1, 0},
		}
	} else if s == "|" {
		diff = []utils.Coord{
			{0, 1},
			{0, -1},
		}
	} else if s == "L" {
		diff = []utils.Coord{
			{0, -1},
			{1, 0},
		}
	} else if s == "F" {
		diff = []utils.Coord{
			{0, 1},
			{1, 0},
		}
	} else if s == "7" {
		diff = []utils.Coord{
			{-1, 0},
			{0, 1},
		}
	} else if s == "J" {
		diff = []utils.Coord{
			{0, -1},
			{-1, 0},
		}
	}
	if len(diff) != 2 {
		panic("Not valid pipe")
	}
	return []utils.Coord{
		{coord.X + diff[0].X, coord.Y + diff[0].Y},
		{coord.X + diff[1].X, coord.Y + diff[1].Y},
	}
}

func parseInput() utils.Coord {
	input := utils.ReadInput("example6.txt")
	m = map[utils.Coord]string{}
	s := utils.Coord{X: -1, Y: -1}
	for y, line := range input {
		for x, r := range line {
			m[utils.Coord{X: x, Y: y}] = utils.ToString(r)
			if r == 'S' {
				s = utils.Coord{X: x, Y: y}
			}
		}
	}
	return s
}
