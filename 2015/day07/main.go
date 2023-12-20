package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/regex"
	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	// day2()
}

func day1() {
	input := utils.ReadInput("example.txt")
	m := map[string]int{}
	for _, line := range input {
		equationParts := strings.Split(line, " -> ")
		key := equationParts[1]

		if strings.Contains(equationParts[0], "AND") {
			parts := strings.Split(equationParts[0], " AND ")
			regex.GetNumberRegexp()

		} else if strings.Contains(equationParts[0], "OR") {

		} else if strings.Contains(equationParts[0], "LSHIFT") {

		} else if strings.Contains(equationParts[0], "RSHIFT") {

		} else if strings.Contains(equationParts[0], "NOT") {

		}

	}
	fmt.Println(input)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func mapToInt(m map[string]int, strings []string) []int {
	for _, str := range strings {

	}
}
