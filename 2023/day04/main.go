package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
	"golang.org/x/exp/slices"
)

type Card struct {
	id      int
	facit   []int
	numbers []int
}

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	cards := parseInput()
	results := []int{}
	for _, card := range cards {
		counter := 0
		for _, number := range card.numbers {
			if slices.Contains(card.facit, number) {
				counter++
			}
		}
		if counter == 0 {
			results = append(results, 0)
		} else {
			sum := 1
			for i := 1; i < counter; i++ {
				sum *= 2
			}
			results = append(results, sum)
		}
	}
	sum := 0
	for _, v := range results {
		sum += v
	}
	println(sum)
}

func day2() {
	originalCards := parseInput()
	winningMap := map[int]int{}
	totalCards := map[int]int{}
	for _, card := range originalCards {
		totalCards[card.id] = 1
		for _, number := range card.numbers {
			if slices.Contains(card.facit, number) {
				winningMap[card.id]++
			}
		}
	}

	for _, card := range originalCards {
		for i := 1; i <= winningMap[card.id]; i++ {
			totalCards[card.id+i] += totalCards[card.id]
		}
	}

	sum := 0
	for _, v := range totalCards {
		sum += v
	}
	fmt.Println(sum)
}

func parseInput() []Card {
	input := utils.ReadInput("dayInput.txt")
	numberRegexp, _ := regexp.Compile("[0-9]+")
	cardList := []Card{}
	for _, cardInput := range input {
		idAndNumbersParts := strings.Split(cardInput, ":")
		idString := numberRegexp.FindAllString(idAndNumbersParts[0], 1)[0]
		facitAndNumbersStrings := strings.Split(idAndNumbersParts[1], "|")
		facitStrings := numberRegexp.FindAllString(facitAndNumbersStrings[0], -1)
		numberStrings := numberRegexp.FindAllString(facitAndNumbersStrings[1], -1)
		card := Card{utils.ToInt(idString), []int{}, []int{}}
		for _, v := range facitStrings {
			card.facit = append(card.facit, utils.ToInt(v))
		}
		for _, v := range numberStrings {
			card.numbers = append(card.numbers, utils.ToInt(v))
		}
		cardList = append(cardList, card)
	}
	return cardList
}
