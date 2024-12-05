package main

import (
	"fmt"
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
	input := utils.ReadInput(EXAMPLE)
	fmt.Println(input)
}

func task2() {
	// input := utils.ReadInput(EXAMPLE)
	// fmt.Println(input)
}
