package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	// day2()
}

func day1() {
	input := parseInput()
	sum := 0
	for _, m := range input {
		horizontal := mirrorHorizontal(m)
		vertical := mirrorVertical(m)
		if vertical >= 0 {
			sum += vertical + 1
		}
		if horizontal >= 0 {
			sum += (horizontal + 1) * 100
		}
	}
	fmt.Println(sum)
}

func mirrorHorizontal(m [][]string) int {
	for i := 0; i < len(m)-1; i++ {
		if bubbleHorizontal(m, i, 0, 0) {
			return i
		}
	}
	return -1
}

func mirrorVertical(m [][]string) int {
	for i := 0; i < len(m[0])-1; i++ {
		if bubbleVertical(m, i, 0, 0) {
			return i
		}
	}
	return -1
}

func bubbleHorizontal(m [][]string, mirrowIndex int, step int, errors int) bool {
	if mirrowIndex-step < 0 || mirrowIndex+step+1 >= len(m) {
		return errors == 1
	}

	for j := 0; j < len(m[mirrowIndex]); j++ {
		if m[mirrowIndex-step][j] != m[mirrowIndex+step+1][j] {
			errors++
		}
		if errors > 1 {
			return false
		}
	}
	return bubbleHorizontal(m, mirrowIndex, step+1, errors)
}

func bubbleVertical(m [][]string, mirrowIndex int, step int, errors int) bool {
	if mirrowIndex-step < 0 || mirrowIndex+step+1 >= len(m[0]) {
		return errors == 1
	}

	for j := 0; j < len(m); j++ {
		if m[j][mirrowIndex-step] != m[j][mirrowIndex+step+1] {
			errors++
		}
		if errors > 1 {
			return false
		}
	}
	return bubbleVertical(m, mirrowIndex, step+1, errors)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func parseInput() [][][]string {
	input := utils.ReadInput("dayInput.txt")
	output := [][][]string{}
	current := [][]string{}
	for _, line := range input {
		if line == "" {
			output = append(output, current)
			current = [][]string{}
			continue
		}
		current = append(current, strings.Split(line, ""))
	}
	output = append(output, current)
	return output
}
