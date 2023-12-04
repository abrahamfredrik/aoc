package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

type card struct {
	id      int
	facit   []int
	numbers []int
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func parseInput() []int {
	input := utils.ReadInput("example.txt")
	numberRegexp, _ := regexp.Compile("[0-9]+")
	for _, cardInput := range input {
		idAndNumbersParts := strings.Split(cardInput, ":")
		idString := numberRegexp.FindAllString(idAndNumbersParts[0], 1)[0]
		facitAndNumbersStrings := strings.Split(idAndNumbersParts[1], "|")
		facitStrings := numberRegexp.FindAllString(facitAndNumbersStrings[0], -1)
		numberStrings := numberRegexp.FindAllString(facitAndNumbersStrings[1], -1)
	}
}
