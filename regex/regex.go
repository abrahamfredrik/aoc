package regex

import "regexp"

func GetNumberRegexp() *regexp.Regexp {
	regex, _ := regexp.Compile("[0-9]+")
	return regex
}
