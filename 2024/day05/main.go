package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/abrahamfredrik/aoc/utils"
)

var EXAMPLE = "exampleInput.txt"
var DAY_INPUT = "dayInput.txt"

func main() {
	start := time.Now()
	task1()
	elapsed := time.Since(start)
	fmt.Println("Time for task 1: ",elapsed)
	fmt.Println("---")
	start2 := time.Now()
	task2()
	elapsed2 := time.Since(start2)
	fmt.Println("Time for task 2: ",elapsed2)
}

func task1() {
	rules, rows := readAndFormatInput()

	correctRows := make([][]int, 0)
	for _, row := range rows {
		isOk := true
		for _, rule := range rules {
			if !rule.isInOrder(row) {
				isOk = false
				break
			}
		}
		if isOk {
			correctRows = append(correctRows, row)
		}
	}
	sum := 0
	for _, row := range correctRows {
		index := (len(row) - 1) / 2
		sum += row[index]
	}
	fmt.Println(sum)
}

func task2() {
	rules, rows := readAndFormatInput()

	invalidRows := make([][]int, 0)
	for _, row := range rows {
		isOk := true
		for _, rule := range rules {
			if !rule.isInOrder(row) {
				isOk = false
				break
			}
		}
		if !isOk {
			invalidRows = append(invalidRows, row)
		}
	}
	correctedInvalidRows := make([][]int, len(invalidRows))

	for rowIndex, row := range invalidRows {
		validRules := findValidRules(rules, row)
		currentSlice := utils.CopySlice(row)
		continueLoop := true
		for continueLoop {
			continueLoop = false
			for _, r := range validRules {
				if !r.isInOrder(currentSlice) {
					continueLoop = true
					temp := utils.CopySlice(currentSlice)
					for i, n := range currentSlice {
						if n == r.first {
							temp[i] = r.second
						} else if n == r.second {
							temp[i] = r.first
						}
					}
					currentSlice = temp
					break
				}
			}
		}
		correctedInvalidRows[rowIndex] = currentSlice
	}

	sum := 0
	for _, row := range correctedInvalidRows {
		index := (len(row) - 1) / 2
		sum += row[index]
	}
	fmt.Println(sum)
}

func findValidRules(rules []rule, numbers []int) []rule {
	validRules := make([]rule, 0)
	for _, rule := range rules {
		if rule.containsNumbers(numbers) {
			validRules = append(validRules, rule)
		}
	}
	return validRules
}

func readAndFormatInput() ([]rule, [][]int) {
	input := utils.ReadInput(DAY_INPUT)
	rules := []rule{}
	rows := make([][]int, 0)
	for _, line := range input {
		if strings.Contains(line, "|") {
			parts := utils.ListToInts(strings.Split(line, "|"))
			rules = append(rules, rule{parts[0], parts[1]})
			if len(parts) != 2 {
				panic("Found strange rule")
			}
		} else if strings.Contains(line, ",") {
			row := utils.ListToInts(strings.Split(line, ","))
			rows = append(rows, row)
		}
	}
	return rules, rows
}
