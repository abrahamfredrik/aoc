package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/abrahamfredrik/aoc/utils"
)

var red string = "red"
var green string = "green"
var blue string = "blue"

func main() {
	day1()
	fmt.Println("---")
	day2()
}

func day1() {
	games := parseInput()

	idSum := 0
	for _, game := range games {
		if game.isGamePosible(12, 13, 14) {
			idSum += game.id
		}
	}
	fmt.Println("Part 1: ", idSum)
}

func day2() {
	games := parseInput()

	power := 0
	for _, game := range games {
		maxGreen, maxRed, maxBlue := math.MinInt, math.MinInt, math.MinInt
		for _, set := range game.sets {
			maxGreen = int(math.Max(float64(set[green]), float64(maxGreen)))
			maxRed = int(math.Max(float64(set[red]), float64(maxRed)))
			maxBlue = int(math.Max(float64(set[blue]), float64(maxBlue)))
		}
		power += maxGreen * maxRed * maxBlue
	}
	fmt.Println("Part 2: ", power)
}

func parseInput() []game {
	input := utils.ReadInput("dayInput.txt")
	games := []game{}
	for _, str := range input {
		games = append(games, createGameFromString(str))
	}
	return games
}

func createGameFromString(str string) game {
	parts := strings.Split(str, ":")
	gameNumber := utils.ToInt(strings.Split(parts[0], " ")[1]) // get number of the game

	sets := []map[string]int{}
	gameSets := strings.Split(parts[1], ";")
	numberRegexp, _ := regexp.Compile("[0-9]+")

	for _, setString := range gameSets {
		m := map[string]int{}
		cubeString := strings.Split(setString, ",")
		for _, s := range cubeString {
			if strings.Contains(s, green) {
				m[green] = utils.ToInt(numberRegexp.FindString(s))
				continue
			}
			if strings.Contains(s, red) {
				m[red] = utils.ToInt(numberRegexp.FindString(s))
				continue
			}
			if strings.Contains(s, blue) {
				m[blue] = utils.ToInt(numberRegexp.FindString(s))
				continue
			}
		}
		sets = append(sets, m)
	}
	return game{gameNumber, sets}
}

type game struct {
	id   int
	sets []map[string]int
}

func (g *game) isGamePosible(redLimit int, greenLimit int, blueLimit int) bool {
	for _, m := range g.sets {
		if m[red] > redLimit || m[green] > greenLimit || m[blue] > blueLimit {
			return false
		}
	}
	return true
}
