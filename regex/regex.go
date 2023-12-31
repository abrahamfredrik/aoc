package regex

import "regexp"

func GetNumberRegexp() *regexp.Regexp {
	regex, _ := regexp.Compile("[0-9]+")
	return regex
}

func GetLowerCaseRegexp() *regexp.Regexp {
	regex, _ := regexp.Compile("[a-z]+")
	return regex
}

func GetTextRegexp() *regexp.Regexp {
	regex, _ := regexp.Compile("[a-zA-Z]+")
	return regex
}
