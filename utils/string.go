package utils

import (
	"sort"
	"strings"
)

func ContainsAllRunes(str string, runes string) bool {
	for _, r := range runes {
		if !strings.ContainsRune(str, r) {
			return false
		}
	}
	return true
}

func SortString(input string) string {
	chars := strings.Split(input, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func IsUpperCase(s string) bool {
	return strings.ToUpper(s) == s
}

func IsLowerCase(s string) bool {
	return strings.ToLower(s) == s
}
