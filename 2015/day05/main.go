package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

var vowels = strings.Split("aeiou", "")
var illegalStrings = []string{
	"ab",
	"cd",
	"pq",
	"xy",
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	input := utils.ReadInput("dayInput.txt")
	niceCounter := 0
	for _, line := range input {
		if !hasDoubleLetter(line) {
			continue
		}
		if !hasVovels(line) {
			continue
		}
		if hasIllegalString(line) {
			continue
		}
		niceCounter++
	}
	fmt.Println(niceCounter)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	niceCounter := 0
	for _, line := range input {
		if !hasOneBetween(line) {
			continue
		}
		if !hasDoublePair(line) {
			continue
		}
		niceCounter++
	}
	fmt.Println(niceCounter)
}

func hasVovels(s string) bool {
	vowelCounter := 0
	for _, letter := range s {
		for _, vowel := range vowels {
			if utils.ToString(letter) == vowel {
				vowelCounter++
			}
		}
	}
	return vowelCounter >= 3
}

func hasDoubleLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}

func hasIllegalString(s string) bool {
	for _, illegalString := range illegalStrings {
		if strings.Contains(s, illegalString) {
			return true
		}
	}
	return false
}

func hasOneBetween(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}

func hasDoublePair(s string) bool {
	for i := 1; i < len(s); i++ {
		tempString := strings.ReplaceAll(s, s[i-1:i+1], "")
		if len(tempString)+4 <= len(s) {
			return true
		}
	}
	return false
}
