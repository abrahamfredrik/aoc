package main

import (
	"fmt"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	mirrorMatrix := parseInput()
	startRay := Ray{utils.Coord{X: -1, Y: 0}, ">"}
	energized := bounceRays(startRay, mirrorMatrix)
	fmt.Println(energized)
	// 6906 correct
}

func day2() {
	mirrorMatrix := parseInput()
	best := 0

	// Right and left
	for i := 0; i < len(mirrorMatrix); i++ {
		startRay := Ray{utils.Coord{X: -1, Y: i}, ">"}
		energized := bounceRays(startRay, mirrorMatrix)
		if energized > best {
			best = energized
		}
		startRay = Ray{utils.Coord{X: len(mirrorMatrix[i]), Y: i}, "<"}
		energized = bounceRays(startRay, mirrorMatrix)
		if energized > best {
			best = energized
		}
	}

	// Top and bottom
	for i := 0; i < len(mirrorMatrix[0]); i++ {
		startRay := Ray{utils.Coord{X: i, Y: -1}, "v"}
		energized := bounceRays(startRay, mirrorMatrix)
		if energized > best {
			best = energized
		}
		startRay = Ray{utils.Coord{X: i, Y: len(mirrorMatrix)}, "^"}
		energized = bounceRays(startRay, mirrorMatrix)
		if energized > best {
			best = energized
		}
	}
	fmt.Println(best)

}

func bounceRays(startRay Ray, mirrorMatrix [][]string) int {
	visited := map[utils.Coord][]string{}
	rays := []Ray{startRay}
	for len(rays) > 0 {
		current := rays[len(rays)-1]
		rays = rays[:len(rays)-1]
		if hasInMap(visited, current) {
			continue
		}

		addToMap(&visited, current)
		if current.direction == ">" {
			c := current.coord.GetRight()
			if c.X >= len(mirrorMatrix[c.Y]) {
				continue
			}
			if mirrorMatrix[c.Y][c.X] == "/" {
				rays = append(rays, Ray{c, "^"})
			} else if mirrorMatrix[c.Y][c.X] == "\\" {
				rays = append(rays, Ray{c, "v"})
			} else if mirrorMatrix[c.Y][c.X] == "|" {
				rays = append(rays, Ray{c, "v"})
				rays = append(rays, Ray{c, "^"})
			} else {
				rays = append(rays, Ray{c, ">"})
			}
		} else if current.direction == "<" {
			c := current.coord.GetLeft()
			if c.X < 0 {
				continue
			}
			if mirrorMatrix[c.Y][c.X] == "/" {
				rays = append(rays, Ray{c, "v"})
			} else if mirrorMatrix[c.Y][c.X] == "\\" {
				rays = append(rays, Ray{c, "^"})
			} else if mirrorMatrix[c.Y][c.X] == "|" {
				rays = append(rays, Ray{c, "v"})
				rays = append(rays, Ray{c, "^"})
			} else {
				rays = append(rays, Ray{c, "<"})
			}
		} else if current.direction == "^" {
			c := current.coord.GetAbove()
			if c.Y < 0 {
				continue
			}
			if mirrorMatrix[c.Y][c.X] == "/" {
				rays = append(rays, Ray{c, ">"})
			} else if mirrorMatrix[c.Y][c.X] == "\\" {
				rays = append(rays, Ray{c, "<"})
			} else if mirrorMatrix[c.Y][c.X] == "-" {
				rays = append(rays, Ray{c, "<"})
				rays = append(rays, Ray{c, ">"})
			} else {
				rays = append(rays, Ray{c, "^"})
			}
		} else if current.direction == "v" {
			c := current.coord.GetBelow()
			if c.Y >= len(mirrorMatrix) {
				continue
			}
			if mirrorMatrix[c.Y][c.X] == "/" {
				rays = append(rays, Ray{c, "<"})
			} else if mirrorMatrix[c.Y][c.X] == "\\" {
				rays = append(rays, Ray{c, ">"})
			} else if mirrorMatrix[c.Y][c.X] == "-" {
				rays = append(rays, Ray{c, "<"})
				rays = append(rays, Ray{c, ">"})
			} else {
				rays = append(rays, Ray{c, "v"})
			}
		} else {
			panic("inavlid direction")
		}
	}
	return len(visited) - 1
}

func parseInput() [][]string {
	input := utils.ReadInput("dayInput.txt")
	m := make([][]string, len(input))
	for y, line := range input {
		m[y] = make([]string, len(line))
		for x, c := range line {
			m[y][x] = utils.ToString(c)
		}
	}
	return m
}

func hasInMap(m map[utils.Coord][]string, ray Ray) bool {
	if val, ok := m[ray.coord]; ok {
		for _, r := range val {
			if ray.direction == r {
				return true
			}
		}
	}
	return false
}

func addToMap(m *map[utils.Coord][]string, ray Ray) {
	if _, ok := (*m)[ray.coord]; !ok {
		(*m)[ray.coord] = []string{}
	}
	(*m)[ray.coord] = append((*m)[ray.coord], ray.direction)
}

type Ray struct {
	coord     utils.Coord
	direction string
}
