package utils

import "fmt"

func PrintMatrix(m [][]string) {
	for _, row := range m {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}
