package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowm := make(map[byte]int)
		colm := make(map[byte]int)
		celm := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				rowm[board[i][j]]++
				if rowm[board[i][j]] > 1 {
					return false
				}
			}
			if board[j][i] >= '1' && board[j][i] <= '9' {
				colm[board[j][i]]++
				if colm[board[j][i]] > 1 {
					return false
				}
			}
			row := (i/3)*3 + j/3
			col := (i%3)*3 + j%3
			if board[row][col] >= '1' && board[row][col] <= '9' {
				celm[board[row][col]]++
				if celm[board[row][col]] > 1 {
					return false
				}
			}
		}
	}

	return true
}
