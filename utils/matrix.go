package utils

import (

)

func CreateMatrix(input []string) [][]string {
	matrix := make([][]string, len(input))
	for r, row := range input {
		matrix[r] = make([]string, len(row))
		for c, col := range row {
			matrix[r][c] = string(col)
		}
	}
	return matrix
}

func CreateIntMatrix(input []string) [][]int {
	matrix := make([][]int, len(input))
	for r, row := range input {
		matrix[r] = make([]int, len(row))
		for c, col := range row {
			matrix[r][c] = int(col)
		}
	}
	return matrix
}