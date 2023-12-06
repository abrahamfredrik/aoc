package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

const (
	ON = iota
	OFF
	TOGGLE
)

type LampAction struct {
	from   utils.Coord
	to     utils.Coord
	action int
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	input := parseInput()
	m := map[[2]int]bool{}
	for _, action := range input {
		for x := action.from.X; x <= action.to.X; x++ {
			for y := action.from.Y; y <= action.to.Y; y++ {
				if action.action == ON {
					m[[2]int{x, y}] = true
				} else if action.action == OFF {
					m[[2]int{x, y}] = false
				} else {
					m[[2]int{x, y}] = !m[[2]int{x, y}]
				}
			}
		}
	}
	counter := 0
	for _, v := range m {
		if v {
			counter++
		}
	}
	fmt.Println(counter)
}

func day2() {
	input := parseInput()
	m := map[[2]int]int{}
	for _, action := range input {
		for x := action.from.X; x <= action.to.X; x++ {
			for y := action.from.Y; y <= action.to.Y; y++ {
				if action.action == ON {
					m[[2]int{x, y}]++
				} else if action.action == OFF {
					if m[[2]int{x, y}] > 0 {
						m[[2]int{x, y}]--
					}
				} else {
					m[[2]int{x, y}] += 2
				}
			}
		}
	}
	counter := 0
	for _, v := range m {
		counter += v
	}
	fmt.Println(counter)
}

func parseInput() []LampAction {
	input := utils.ReadInput("dayInput.txt")
	coordRegex, _ := regexp.Compile("[0-9]+,[0-9]+")
	lampActions := []LampAction{}
	for _, line := range input {
		action := findAction(line)
		coordStrings := coordRegex.FindAllString(line, -1)
		from := utils.ListToInts(strings.Split(coordStrings[0], ","))
		to := utils.ListToInts(strings.Split(coordStrings[1], ","))
		lampActions = append(lampActions, LampAction{utils.Coord{X: from[0], Y: from[1]}, utils.Coord{X: to[0], Y: to[1]}, action})
	}
	return lampActions
}

func findAction(s string) int {
	if strings.Contains(s, "turn on") {
		return ON
	} else if strings.Contains(s, "turn off") {
		return OFF
	} else if strings.Contains(s, "toggle") {
		return TOGGLE
	} else {
		panic("No matching action")
	}
}
