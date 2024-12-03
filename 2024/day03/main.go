package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

var mulRegex, _ = regexp.Compile(`mul\(\d*,\d*\)`)

func day1() {
	input := strings.Join(utils.ReadInput("dayInput.txt"), "")
	countMultiplication(input)
}

func day2() {
	input := strings.Join(utils.ReadInput("dayInput.txt"), "")

	cleanedInput := ""
	buffer := ""
	active := true
	for _, s := range input {
		buffer = buffer + string(s)
		if active {
			cleanedInput = cleanedInput + string(s)
		}

		if strings.HasSuffix(buffer, "do()") {
			active = true
		} else if strings.HasSuffix(buffer, "don't()") {
			active = false
		}
	}

	countMultiplication(cleanedInput)
}

func countMultiplication(input string) {
	hits := mulRegex.FindAllString(input, -1)
	sum := 0
	for _, hit := range hits {
		firstRemoved := strings.ReplaceAll(hit, `mul(`, ``)
		allRemoved := strings.ReplaceAll(firstRemoved, `)`, ``)
		numbersString := strings.Split(allRemoved, ",")
		numbers := utils.ListToInts(numbersString)
		sum += numbers[0] * numbers[1]
	}
	fmt.Println(sum)
}
