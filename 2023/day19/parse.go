package main

import (
	"strings"

	"github.com/abrahamfredrik/aoc/regex"
	"github.com/abrahamfredrik/aoc/utils"
)

func parseInput() {
	input := utils.ReadInput("example.txt")

	workFlows := []workFlow{}
	parts := []part{}
	parsParts := false
	for _, line := range input {

	}

}

func parsePart(s string) part {
	part := part{}
	stringParts := strings.Split(s, ",")
	for _, p := range stringParts {
		if strings.Contains(p, "x") {
			part.x = utils.ToInt(regex.GetNumberRegexp().FindString(p))

		} else if strings.Contains(p, "m") {
			part.m = utils.ToInt(regex.GetNumberRegexp().FindString(p))

		} else if strings.Contains(p, "a") {
			part.a = utils.ToInt(regex.GetNumberRegexp().FindString(p))

		} else if strings.Contains(p, "s") {
			part.s = utils.ToInt(regex.GetNumberRegexp().FindString(p))

		} else {
			panic("invalid string")
		}
	}
	return part
}

func parseWorkFlow(s string) workFlow {

}
