package main

import (
	"fmt"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	input := utils.ReadInput("dayInput.txt")
	counter := 0
	for _, line := range input {
		for _, r := range line {
			if r == '(' {
				counter++
			} else {
				counter--
			}
		}
	}
	fmt.Println(counter)
}

func day2() {
	input := utils.ReadInput("dayInput.txt")
	counter := 0
	index := 1
	for _, line := range input {
		for _, r := range line {
			if r == '(' {
				counter++
			} else {
				counter--
			}
			if counter<0{
				fmt.Println(index)
				return
			}
			index++
		}
	}
	
}
