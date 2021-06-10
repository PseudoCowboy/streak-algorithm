package main

func exist(board [][]byte, word string) bool {
	for i := range board {
			for j := range board[0] {
					used := make([][]bool, len(board))
					for k := range used {
							used[k] = make([]bool, len(board[0]))
					}
					result := false
					search(board, word, 0, i, j, &used, &result)
					if result {
							return true
					}
			}
	}

	return false
}

var move = [][]int{
	{0,1},
	{0,-1},
	{1,0},
	{-1,0},
}

func search(board [][]byte, word string, index, x, y int, used *[][]bool, result *bool) {
	if board[x][y] != word[index] {
			return
	}
	index++
	if index == len(word) {
			*result = true
			return
	}

	(*used)[x][y] = true
	for i := range move {
			nextX := x + move[i][0]
			nextY := y + move[i][1]
			if nextX >= 0 && nextX < len(board) && nextY >= 0 && nextY < len(board[0]) && !(*used)[nextX][nextY] {
					search(board, word, index, nextX, nextY, used, result)
					(*used)[nextX][nextY] = false
			}
	}
}
