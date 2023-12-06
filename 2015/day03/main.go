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
	input := utils.ReadInput("dayInput.txt")
	current := [2]int{0, 0}
	m := map[[2]int]int{current: 1}
	for _, line := range input {
		for _, direction := range line {
			if direction == '<' {
				current = [2]int{current[0] - 1, current[1]}
			} else if direction == '>' {
				current = [2]int{current[0] + 1, current[1]}
			} else if direction == '^' {
				current = [2]int{current[0], current[1] + 1}
			} else if direction == 'v' {
				current = [2]int{current[0], current[1] - 1}
			} else {
				panic("Not valid arrow")
			}
			m[current]++
		}
	}

	fmt.Println(len(m))
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	currentSanta := [2]int{0, 0}
	currentRobo := [2]int{0, 0}
	m := map[[2]int]int{currentSanta: 1}
	santasTurn := true
	for _, line := range input {
		for _, direction := range line {
			current := currentSanta
			if !santasTurn {
				current = currentRobo
			}
			if direction == '<' {
				current = [2]int{current[0] - 1, current[1]}
			} else if direction == '>' {
				current = [2]int{current[0] + 1, current[1]}
			} else if direction == '^' {
				current = [2]int{current[0], current[1] + 1}
			} else if direction == 'v' {
				current = [2]int{current[0], current[1] - 1}
			} else {
				panic("Not valid arrow")
			}
			m[current]++
			if santasTurn {
				currentSanta = current
			} else {
				currentRobo = current
			}
			santasTurn = !santasTurn
		}
	}

	fmt.Println(len(m))
}
