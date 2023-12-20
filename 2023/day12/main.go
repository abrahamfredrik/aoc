package main

import (
	"fmt"
	"regexp"
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
	input := parseInput()
	for _, blueprint := range input {
		fmt.Println(blueprint)
	}

}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func parseInput() []Blueprint {
	input := utils.ReadInput("example.txt")
	springRegex, _ := regexp.Compile("(\\.|#|\\?)+")
	blueprints := []Blueprint{}
	for _, line := range input {
		springLine := springRegex.FindAllString(line, -1)
		if len(springLine) > 1 {
			panic("condition not unique")
		}
		conditions := regex.GetNumberRegexp().FindAllString(line, -1)
		blueprints = append(blueprints, Blueprint{strings.Split(springLine[0], ""), utils.ListToInts(conditions)})
	}
	return blueprints
}

type Blueprint struct {
	springs    []string
	conditions []int
}
