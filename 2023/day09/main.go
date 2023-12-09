package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
	"golang.org/x/exp/slices"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	part1Func := func(history [][]int) int {
		historyIndex := len(history)
		sum := 0
		for historyIndex > 0 {
			historyIndex--
			size := len(history[historyIndex])
			sum += history[historyIndex][size-1]
		}
		return sum
	}
	solution(part1Func)
}

func day2() {
	part2Func := func(history [][]int) int {
		for _, line := range history {
			slices.Reverse(line) // Reverse to make handlnig easier
		}
		historyIndex := len(history)
		sum := 0
		for historyIndex > 0 {
			historyIndex--
			size := len(history[historyIndex])
			sum = history[historyIndex][size-1] - sum
		}
		return sum
	}
	solution(part2Func)
}

func solution(function func([][]int) int) {
	input := parseInput()
	extrapolatedValues := make([]int, len(input))
	for rowIndex, row := range input {
		history := createDiffHistory(row)
		sum := function(history)
		extrapolatedValues[rowIndex] = sum
	}
	power := 0
	for _, v := range extrapolatedValues {
		power += v
	}
	fmt.Println(power)
}

func createDiffHistory(row []int) [][]int {
	history := [][]int{row}
	historyIndex := 0
	for !isAllZero(history[historyIndex]) {
		workingList := []int{}
		for i := 0; i < len(history[historyIndex])-1; i++ {
			diff := history[historyIndex][i+1] - history[historyIndex][i]
			workingList = append(workingList, diff)
		}
		history = append(history, workingList)
		historyIndex++
	}
	return history
}

func isAllZero(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

func parseInput() [][]int {
	input := utils.ReadInput("dayInput.txt")
	intList := make([][]int, len(input))
	for i, line := range input {
		intList[i] = utils.ListToInts(strings.Split(line, " "))
	}
	return intList
}
