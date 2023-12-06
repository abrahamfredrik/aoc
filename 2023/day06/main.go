package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/regex"
	"github.com/abrahamfredrik/aoc/utils"
)

type TimeDistance struct {
	time     int
	distance int
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	timeDistanceList := parseInput()
	solution(timeDistanceList)
}

func day2() {
	timeDistanceList := parseInput2()
	solution(timeDistanceList)
}

func solution(timeDistanceList []TimeDistance){
	options := []int{}
	for _, timeDist := range timeDistanceList {
		counter := 0
		for time := 0; time <= timeDist.time; time++ {
			traveledDistance := time * (timeDist.time - time)
			if traveledDistance > timeDist.distance {
				counter++
			}
		}
		options = append(options, counter)
	}
	total := 1
	for _, v := range options {
		total *= v
	}
	fmt.Println(total)

}

func parseInput() []TimeDistance {
	input := utils.ReadInput("dayInput.txt")
	numberRegexp := regex.GetNumberRegexp()
	timeDistanceList := []TimeDistance{}
	time := utils.ListToInts(numberRegexp.FindAllString(input[0], -1))
	distance := utils.ListToInts(numberRegexp.FindAllString(input[1], -1))
	if len(time) != len(distance) {
		panic("Not same number time and distance")
	}
	for i, v := range time {
		timeDistanceList = append(timeDistanceList, TimeDistance{v, distance[i]})
	}
	return timeDistanceList
}

func parseInput2() []TimeDistance {
	input := utils.ReadInput("dayInput.txt")
	numberRegexp := regex.GetNumberRegexp()

	time := utils.ToInt(strings.Join(numberRegexp.FindAllString(input[0], -1), ""))
	distance := utils.ToInt(strings.Join(numberRegexp.FindAllString(input[1], -1), ""))
	return []TimeDistance{{time, distance}}
}
