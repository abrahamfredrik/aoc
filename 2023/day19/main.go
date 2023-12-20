package main

import (
	"fmt"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	// day2()
}

func day1() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	fmt.Println(input)
}



type workFlow struct {
	key string
	exp string
}

type part struct {
	x int
	m int
	a int
	s int
}
