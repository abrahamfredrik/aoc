package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
	"github.com/abrahamfredrik/aoc/utils/math"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	goalList := pathFinder([]string{"AAA"})
	fmt.Println(goalList)
}

func day2() {
	_, m := parseInput()
	positionList := []string{}
	for k := range m {
		if k[2] == 'A' {
			positionList = append(positionList, k)
		}
	}
	goalList := pathFinder(positionList)

	fmt.Println("LCM: ", aocMath.LeastCommonMultiple(goalList[0], goalList[1], goalList[2:]...))
	for _, v := range goalList {
		fmt.Println(v)
	}
}

func pathFinder(startingPoints []string) []int {
	instructions, m := parseInput()
	stepsToGoal := make([]int, len(startingPoints))
	for i, start := range startingPoints {
		index := 0
		counter := 0

		for true {
			inst := instructions[index]
			if inst == "R" {
				start = m[start].right
			} else {
				start = m[start].left
			}
			counter++
			if start[2] == 'Z' {
				break
			}

			if index == len(instructions)-1 {
				index = 0
			} else {
				index++
			}
		}
		stepsToGoal[i] = counter
	}
	return stepsToGoal
}

func parseInput() ([]string, map[string]element) {
	input := utils.ReadInput("dayInput.txt")
	inst := strings.Split(input[0], "")
	elementRegexp, _ := regexp.Compile("[0-9A-Z]{3}")
	m := map[string]element{}
	for _, el := range input[2:] {
		parts := elementRegexp.FindAllString(el, -1)
		if len(parts) != 3 {
			panic("Fail in parsing")
		}
		m[parts[0]] = element{parts[1], parts[2]}
	}
	return inst, m
}

type element struct {
	left  string
	right string
}
