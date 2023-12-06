package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/abrahamfredrik/aoc/utils"
)

type AlmanacRange struct {
	destStart   int
	sourceStart int
	rangeLength int
}

var mapStrings = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	seeds, inputMap := parseInput()
	start := time.Now()
	min := math.MaxInt
	for _, seed := range seeds {
		for _, keyString := range mapStrings {
			seed = mapNumber(seed, inputMap[keyString])
		}
		if min > seed {
			min = seed
		}
	}
	fmt.Println(min)
	fmt.Println("Time: ", time.Since(start))
}

func day2() { // Could be smarter but hey, bruteforce solves it in 5-6 seconds so why bother?
	seeds, inputMap := parseInput()
	current := 0
	start := time.Now()
	for true {
		temp := current
		for i := len(mapStrings) - 1; i >= 0; i-- {
			temp = reverseMapNumber(temp, inputMap[mapStrings[i]])
		}
		if hasSeed(seeds, temp) {
			break
		}
		current++
	}
	fmt.Println(current)
	fmt.Println("Time: ", time.Since(start))
}

func hasSeed(seeds []int, seed int) bool {
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		seedRange := seeds[i+1]
		if seed >= start && seed <= start+seedRange {
			return true
		}
	}
	return false
}

func mapNumber(number int, mapper []AlmanacRange) int {
	for _, aRange := range mapper {
		diff := number - aRange.sourceStart
		if diff >= 0 && diff <= aRange.rangeLength {
			return aRange.destStart + diff
		}
	}
	return number
}

func reverseMapNumber(number int, mapper []AlmanacRange) int {
	for _, aRange := range mapper {
		diff := number - aRange.destStart
		if diff >= 0 && diff <= aRange.rangeLength {
			return aRange.sourceStart + diff
		}
	}
	return number
}

func parseInput() ([]int, map[string][]AlmanacRange) {
	input := utils.ReadInput("dayInput.txt")
	numberRegexp, _ := regexp.Compile("[0-9]+")
	seeds := utils.ListToInts(numberRegexp.FindAllString(input[0], -1))

	m := map[string][]AlmanacRange{}
	currentMap := ""
	for _, line := range input[2:] {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-to-") {
			currentMap = strings.Split(line, " ")[0]
			if m[currentMap] == nil {
				m[currentMap] = []AlmanacRange{}
			}
			continue
		}
		numbers := utils.ListToInts(numberRegexp.FindAllString(line, -1))
		if len(numbers) != 3 {
			panic("Wrong length on input")
		}
		m[currentMap] = append(m[currentMap], AlmanacRange{numbers[0], numbers[1], numbers[2]})
	}
	return seeds, m
}
