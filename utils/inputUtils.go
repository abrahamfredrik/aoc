package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineList []string
	for scanner.Scan() {
		line := scanner.Text()
		lineList = append(lineList, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lineList
}
