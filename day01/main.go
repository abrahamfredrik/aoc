package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/abrahamfredrik/aoc_2021_go/utils"
)

func main() {
	input := utils.ReadInput("example2.txt")
	sum := 0
	for _, row := range input {
		letters := []rune{}
		replacedRow := replaceNumbers(row)
		fmt.Println(row, "-", replacedRow)
		for _, r := range replacedRow {
			if unicode.IsDigit(r) {
				letters = append(letters, r)
			}
		}
		number := utils.ToString(letters[0]) + utils.ToString(letters[len(letters)-1])
		sum += utils.ToInt(number)
	}
	fmt.Println(sum)
}

func replaceNumbers(s string) string {
	
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
