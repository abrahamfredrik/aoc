package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/abrahamfredrik/aoc/utils"
)

func main() {
	solution("00000")
	fmt.Println("---")
	solution("000000")
}

func solution(hashStartsWith string) {
	input := "bgvyzdsv"
	hex := ""
	i := 0
	for true {
		hex = GetMD5Hash(input + utils.ToString(i))
		if hex[0:len(hashStartsWith)] == hashStartsWith {
			break
		}
		i++
	}
	fmt.Println(i)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
