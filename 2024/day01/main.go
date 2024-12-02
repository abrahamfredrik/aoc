package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	left, right := setupList()
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i, value := range left {
		sum += int(math.Abs(float64(value) - float64(right[i])))
	}
	fmt.Println(sum)
}

func day2() {
	left, right := setupList()
	sum := 0
	for _, leftValue := range left {
		for _, rightValue := range right {
			if leftValue == rightValue {
				sum += leftValue
			}
		}
		
	}
	fmt.Println(sum)
}

func setupList() ([]int,[]int) {
	input := utils.ReadInput("dayInput.txt")
	left := make([]int, len(input))
	right := make([]int, len(input))
	for i, row := range input {
		fields := strings.Fields(row)
		left[i] = utils.ToInt(fields[0])
		right[i] = utils.ToInt(fields[1])
	}
	return left, right
}
