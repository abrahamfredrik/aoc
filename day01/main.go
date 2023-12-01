package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	day2()
}

func day1() {
	input := utils.ReadInput("dayInput.txt")
	sum := 0
	for _, row := range input {
		numbers := []string{}
		for _, r := range row {
			if unicode.IsDigit(r) {
				numbers = append(numbers, utils.ToString(r))
			}
		}

		sum += utils.ToInt(numbers[0] + numbers[len(numbers)-1])
	}
	fmt.Println("Day 1: ", sum)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	sum := 0
	for _, row := range input {
		number := findFirstAndLast(row)
		sum += utils.ToInt(number)
	}
	fmt.Println("Day 2: ", sum)
}

func findFirstAndLast(s string) int {
	firstString := ""
	firstDigit := ""
firstOuter:
	for _, r := range s {
		firstString = firstString + utils.ToString(r)
		firstString = replaceNumber(firstString)
		for _, fr := range firstString {
			if unicode.IsDigit(fr) {
				firstDigit = utils.ToString(fr)
				break firstOuter
			}
		}
	}
	lastString := ""
	lastDigit := ""
lastOuter:
	for i := len(s) - 1; i >= 0; i-- {
		lastString = utils.ToString(s[i]) + lastString
		lastString = replaceNumber(lastString)
		for _, fr := range lastString {
			if unicode.IsDigit(fr) {
				lastDigit = utils.ToString(fr)
				break lastOuter
			}
		}
	}

	return utils.ToInt(firstDigit + lastDigit)
}

func replaceNumber(s string) string {
	// Maybe not the prettiest but works
	s = strings.ReplaceAll(s, "nine", "9")
	s = strings.ReplaceAll(s, "eight", "8")
	s = strings.ReplaceAll(s, "seven", "7")
	s = strings.ReplaceAll(s, "six", "6")
	s = strings.ReplaceAll(s, "five", "5")
	s = strings.ReplaceAll(s, "four", "4")
	s = strings.ReplaceAll(s, "three", "3")
	s = strings.ReplaceAll(s, "two", "2")
	s = strings.ReplaceAll(s, "one", "1")
	return s
}
