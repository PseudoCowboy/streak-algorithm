package main

func updateMatrix(mat [][]int) [][]int {
	ans := make([][]int, len(mat))
	for i := range ans {
		ans[i] = make([]int, len(mat[0]))
	}
	for i := range ans {
		for j := range ans[i] {
			if mat[i][j] != 0 {
				ans[i][j] = bfs(mat, i, j)
			}
		}
	}
	return ans
}

var move = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func bfs(mat [][]int, x, y int) int {
	used := make([][]bool, len(mat))
	for i := range used {
		used[i] = make([]bool, len(mat[0]))
	}
	used[x][y] = true
	stack := make([][]int, 0)
	depth := 0
	stack = append(stack, []int{x, y})
	for len(stack) != 0 {
		n := len(stack)
		for i := range stack {
			if mat[stack[i][0]][stack[i][1]] == 0 {
				return depth
			}
			for j := range move {
				nextX := stack[i][0] + move[j][0]
				nextY := stack[i][1] + move[j][1]
				if nextX >= 0 && nextY >= 0 && nextX < len(mat) && nextY < len(mat[0]) && !used[nextX][nextY] {
					used[nextX][nextY] = true
					stack = append(stack, []int{nextX, nextY})
				}
			}
		}
		depth++
		stack = stack[n:]
	}
	return depth
}

func updateMatrix1(mat [][]int) [][]int {
	ans := make([][]int, len(mat))
	q := make([][]int, 0)
	for i := range ans {
		ans[i] = make([]int, len(mat[0]))
		for j := range ans[i] {
			if mat[i][j] == 0 {
				ans[i][j] = -1
				q = append(q, []int{i, j})
			}
		}
	}
	level := 1
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			size--
			node := q[0]
			q = q[1:]
			for _, m := range move {
				nextX := node[0] + m[0]
				nextY := node[1] + m[1]
				if nextX < 0 || nextY < 0 || nextX >= len(mat) || nextY >= len(mat[0]) || ans[nextX][nextY] > 0 || ans[nextX][nextY] < 0 {
					continue
				}
				ans[nextX][nextY] = level
				q = append(q, []int{nextX, nextY})
			}
		}
		level++
	}
	for i := range ans {
		for j := range ans[i] {
			if ans[i][j] == -1 {
				ans[i][j] = 0
			}
		}
	}

	return ans
}
