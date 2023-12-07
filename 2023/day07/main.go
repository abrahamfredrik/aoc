package main

import (
	"fmt"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
	"golang.org/x/exp/slices"
)

func main() {
	day2() // Only day 2 solution is printed
}

func day2() {
	listOfHands := parseInput()
	for i := 0; i < len(listOfHands); i++ {
		calculatePowerOfHand(&listOfHands[i])
	}
	slices.SortFunc(listOfHands, compareHandsByPower)
	rank := 1
	sum := 0
	for _, hand := range listOfHands {
		fmt.Println(hand)
		sum += rank * hand.bid
		rank++
	}
	fmt.Println(sum)
}

func compareHandsByPower(a handOfCards, b handOfCards) int {
	return a.handValue - b.handValue
}

func calculatePowerOfHand(hand *handOfCards) { // Calculating the power makes it easy to compare
	factor := 10000000000
	if isFiveOfKind(hand.cards) {
		hand.handValue += factor * 9
	} else if isFourOfKind(hand.cards) {
		hand.handValue += factor * 8
	} else if isFullhouse(hand.cards) {
		hand.handValue += factor * 7
	} else if isThreeOfKind(hand.cards) {
		hand.handValue += factor * 6
	} else if isTwoPair(hand.cards) {
		hand.handValue += factor * 5
	} else if isPair(hand.cards) {
		hand.handValue += factor * 4
	}
	factor = factor / 100
	for _, card := range hand.cards {
		hand.handValue += factor * card
		factor = factor / 100
	}
}

func parseInput() []handOfCards {
	input := utils.ReadInput("dayInput.txt")
	listOfHands := []handOfCards{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		bid := utils.ToInt(parts[1])
		hand := []int{}
		for _, s := range strings.Split(parts[0], "") {
			if s == "T" {
				hand = append(hand, 10)
			} else if s == "J" {
				hand = append(hand, 1) // Joker least worth
			} else if s == "Q" {
				hand = append(hand, 12)
			} else if s == "K" {
				hand = append(hand, 13)
			} else if s == "A" {
				hand = append(hand, 14)
			} else {
				hand = append(hand, utils.ToInt(s))
			}
		}
		listOfHands = append(listOfHands, handOfCards{bid, hand, 0})
	}
	return listOfHands
}

func isFiveOfKind(hand []int) bool {
	return isXOfKind(hand, 5)
}

func isFourOfKind(hand []int) bool {
	return isXOfKind(hand, 4)
}

func isThreeOfKind(hand []int) bool {
	return isXOfKind(hand, 3) && !isFullhouse(hand)
}

func isXOfKind(hand []int, x int) bool {
	m := map[int]int{}
	for _, c := range hand {
		m[c]++
	}
	for k, v := range m {
		if (k != 1 && v+m[1] == x) || m[1] == x {
			return true
		}
	}
	return false
}

func isFullhouse(hand []int) bool {
	m := map[int]int{}
	for _, c := range hand {
		m[c]++
	}
	if m[1] == 0 {
		has2 := false
		has3 := false
		for _, v := range m {
			if v == 3 {
				has3 = true
			} else if v == 2 {
				has2 = true
			}
		}
		return has2 && has3
	}
	counter := 0
	for _, v := range m {
		if v == 2 {
			counter++
		}
	}
	return counter == 2
}

func isTwoPair(hand []int) bool {
	return countPairs(hand) == 2
}

func isPair(hand []int) bool {
	return countPairs(hand) == 1 && !isFullhouse(hand)
}

func countPairs(hand []int) int {
	m := map[int]int{}
	for _, c := range hand {
		m[c]++
	}
	counter := m[1]
	for _, v := range m {
		if v == 2 {
			counter++
		}
	}
	return counter
}

type handOfCards struct {
	bid       int
	cards     []int
	handValue int
}
