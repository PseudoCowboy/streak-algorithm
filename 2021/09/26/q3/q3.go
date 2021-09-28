package main

import "fmt"

func main() {
	fmt.Println(placeWordInCrossword([][]byte{
		{' ', '#', 'a'},
		{' ', '#', 'c'},
		{' ', '#', 'a'},
	}, "ac"))
}

func placeWordInCrossword(board [][]byte, word string) bool {
	n := len(word)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '#' {
				continue
			}
			if board[i][j] != ' ' && board[i][j] != word[0] {
				continue
			}
			strList := find(board, i, j, n)
			fmt.Println(strList)
			for _, str := range strList {
				ans := check(str, word)
				if ans {
					return ans
				}
			}
		}
	}
	return false
}

func check(str, word string) bool {
	for i := range str {
		if str[i] == ' ' {
			continue
		}
		if str[i] != word[i] {
			return false
		}
	}
	return true
}

func find(board [][]byte, r, c, n int) []string {
	ans := make([]string, 0)
	for _, move := range moves {
		current := make([]byte, 0)
		current = append(current, board[r][c])
		nextr := r + move[0]
		nextc := c + move[1]
		if r-move[0] < len(board) && r-move[0] >= 0 && c-move[1] >= 0 && c-move[1] < len(board[0]) && board[r-move[0]][c-move[1]] != '#' {
			continue
		}
		for nextc >= 0 && nextc < len(board[0]) && nextr >= 0 && nextr < len(board) && board[nextr][nextc] != '#' {
			current = append(current, board[nextr][nextc])
			nextr += move[0]
			nextc += move[1]
		}
		if len(string(current)) != n {
			continue
		} else {
			fmt.Println(len(current), n, string(current))
			ans = append(ans, string(current))
		}
	}
	return ans
}

var moves = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}
