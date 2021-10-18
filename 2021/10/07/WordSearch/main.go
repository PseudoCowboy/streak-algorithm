package main

func main() {
	exist([][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}, "SEE")
}

var moves = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			used[i][j] = true
			if word[0] == board[i][j] && search(board, word, string(board[i][j]), i, j, used) {
				return true
			}
			used[i][j] = false
		}
	}
	return false
}

func search(board [][]byte, word, current string, row, col int, used [][]bool) bool {
	if current == word {
		return true
	}

	for _, v := range moves {
		nr := row + v[0]
		nc := col + v[1]
		if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) && !used[nr][nc] && board[nr][nc] == word[len(current)] {
			current += string(board[nr][nc])
			used[nr][nc] = true
			if search(board, word, current, nr, nc, used) {
				return true
			}
			current = current[:len(current)-1]
			used[nr][nc] = false
		}

	}
	return false
}
