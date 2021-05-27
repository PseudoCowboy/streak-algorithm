package main

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				search(grid, &visited, i, j)
				res++
			}
		}
	}

	return res
}

func search(grid [][]byte, visited *[][]bool, x, y int) {
	(*visited)[x][y] = true
	for i := 0; i < len(dir); i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(grid, nx, ny) && !(*visited)[nx][ny] && grid[nx][ny] == '1' {
			search(grid, visited, nx, ny)
		}
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
