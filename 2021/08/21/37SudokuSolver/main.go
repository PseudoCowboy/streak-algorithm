package main

import "fmt"

func main() {
	solveSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
}

func solveSudoku(board [][]byte) {
	dfs(board)
	fmt.Println(board)
}

func dfs(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			var k byte
			for k = '1'; k <= '9'; k++ {
				if check(board, i, j, k) {
					board[i][j] = k
					if dfs(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func check(board [][]byte, x, y int, val byte) bool {
	srow, scol := (x/3)*3, (y/3)*3
	for i := 0; i < 9; i++ {
		if board[x][i] == val {
			return false
		}
		if board[i][y] == val {
			return false
		}
		if board[srow+i%3][scol+i/3] == val {
			return false
		}
	}
	return true
}
