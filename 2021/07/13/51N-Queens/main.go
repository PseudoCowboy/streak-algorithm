package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	current := make([][]string, n)
	for i := range current {
		temp := make([]string, n)
		for j := range temp {
			temp[j] = "."
		}
		current[i] = temp
	}
	dfs(&result, current, 0)
	return result
}

func dfs(result *[][]string, current [][]string, row int) {
	if row == len(current) {
		temp := make([]string, len(current))
		for i := range temp {
			temp[i] = strings.Join(current[i], "")
		}
		*result = append(*result, temp)
		return
	}

	for i := range current[row] {
		if isValid(current, row, i) {
			current[row][i] = "Q"
			dfs(result, current, row+1)
			current[row][i] = "."
		}
	}
}

func isValid(current [][]string, row, col int) bool {
	for i := row; i >= 0; i-- {
		if current[i][col] == "Q" {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if current[i][j] == "Q" {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j < len(current[0]); i, j = i-1, j+1 {
		if current[i][j] == "Q" {
			return false
		}
	}

	return true
}
