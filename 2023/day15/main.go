package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/regex"
	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	// day1()
	fmt.Println("---")
	day2()
}

func day1() {
	words := parseInput()
	sum := 0
	for _, word := range words {
		sum += hashWord(word)
	}
	fmt.Println(sum)
}

func day2() {
	words := parseInput()
	boxes := map[int][]Lens{}
	for i := 0; i < 256; i++ {
		boxes[i] = []Lens{}
	}

	for _, word := range words {
		lens := createLens(word)
		hash := hashWord(lens.name)
		if lens.operator == '=' {
			exist := false
			for i, l := range boxes[hash] {
				if l.name == lens.name {
					boxes[hash][i].value = lens.value
					exist = true
					break
				}
			}
			if !exist {
				boxes[hash] = append(boxes[hash], lens)
			}
		} else {
			for _, l := range boxes[hash] {
				if l.name == lens.name {
					boxes[hash] = removeFromSlice(boxes[hash], l)
					break
				}
			}
		}
	}

	power := 0
	for k, v := range boxes {
		box := k + 1
		for i, lens := range v {
			power += box * (i + 1) * lens.value
		}
	}
	fmt.Println(power)
}

func createLens(s string) Lens {
	numberOptional := regex.GetNumberRegexp().FindAllString(s, -1)
	var value int
	if len(numberOptional) > 0 {
		value = utils.ToInt(numberOptional[0])
	} else {
		value = 0
	}
	name := regex.GetTextRegexp().FindAllString(s, -1)[0]
	operatorRegex, _ := regexp.Compile("=|-")
	operator := operatorRegex.FindAllString(s, -1)[0][0]
	return Lens{name, rune(operator), value}
}

func hashWord(s string) int {
	value := 0
	for _, r := range s {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}

func removeFromSlice(s []Lens, remove Lens) []Lens {
	for i, v := range s {
		if v == remove {
			return append(s[:i], s[i+1:]...)
		}
	}
	panic(fmt.Sprintf("Could not remove %v from slice %v", remove, s))
}

func parseInput() []string {
	input := utils.ReadInput("dayInput.txt")
	return strings.Split(input[0], ",")
}

type Lens struct {
	name     string
	operator rune
	value    int
}
