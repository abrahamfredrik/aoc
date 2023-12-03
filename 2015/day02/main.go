package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
	"golang.org/x/exp/slices"
)

type box struct {
	length int
	height int
	width  int
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	boxes := parseInput()
	total := 0
	for _, box := range boxes {
		side1 := box.height * box.length
		side2 := box.height * box.width
		side3 := box.length * box.width
		total += 2*(side1+side2+side3) + int(math.Min(math.Min(float64(side1), float64(side2)), float64(side3)))
	}
	fmt.Println(total)
}

func day2() {
	boxes := parseInput()
	total := 0
	for _, box := range boxes {
		sides := []int{box.height, box.length, box.width}
		slices.Sort(sides)
		total += sides[0] + sides[0] + sides[1] + sides[1] + sides[0]*sides[1]*sides[2]
	}
	fmt.Println(total)
}

func parseInput() []box {
	input := utils.ReadInput("dayInput.txt")
	boxes := []box{}
	for _, line := range input {
		parts := strings.Split(line, "x")
		boxes = append(boxes, box{utils.ToInt(parts[0]), utils.ToInt(parts[1]), utils.ToInt(parts[2])})
	}
	return boxes
}
