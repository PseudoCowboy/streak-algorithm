package main

func solveSudoku(board [][]byte) {
	dfs(board)
}

func dfs(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				continue
			}
			var k byte
			for k = '1'; k <= '9'; k++ {
				if isValid(row, col, k, board) {
					board[row][col] = k
					if dfs(board) {
						return true
					}
					board[row][col] = '.'
				}
			}
			return false
		}
	}

	return true
}

func isValid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
