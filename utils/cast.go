package utils

import (
	"fmt"
	"log"
	"strconv"
)

func ToString(input interface{}) string {
	switch input.(type) {
	case int:
		return strconv.Itoa(input.(int))
	case rune:
		return string(input.(rune))
	case byte:
		return string(rune(input.(byte)))
	default:
		panic(fmt.Sprintf("Unhandled type for string casting %T", input))
	}
}

func ListToString[T int | rune | byte](list []T) []string {
	strList := make([]string, len(list))
	for i, v := range list {
		strList[i] = ToString(v)
	}
	return strList
}

func ToInt(input interface{}) int {
	switch input.(type) {
	case int:
		return input.(int)
	case string:
		val, err := strconv.Atoi(input.(string))
		if err != nil {
			panic("Error converting string to int " + err.Error())
		}
		return val
	case rune:
		return (int(input.(rune)) - '0')
	default:
		panic(fmt.Sprintf("Unhandled type for int casting %T", input))
	}
}

func ListToInts(list []string) []int {
	intList := make([]int, len(list))
	for i, v := range list {
		intList[i] = ToInt(v)
	}
	return intList
}

func BinaryToDecimal(str string) int {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
