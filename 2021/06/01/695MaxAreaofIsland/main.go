package main

var move = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			start := 0
			search(grid, &start, &max, i, j)
		}
	}
	return max
}

func search(grid [][]int, current *int, max *int, x, y int) {
	if grid[x][y] == 1 {
		grid[x][y] = 0
		*current += 1
		if *current > *max {
			*max = *current
		}
	} else {
		return
	}

	for i := range move {
		nextX := x + move[i][0]
		nextY := y + move[i][1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) {
			search(grid, current, max, nextX, nextY)
		}
	}
}

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func maxAreaOfIsland1(grid [][]int) int {
	res := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}
			area := areaOfIsland(grid, i, j)
			if area > res {
				res = area
			}
		}
	}
	return res
}

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func areaOfIsland(grid [][]int, x, y int) int {
	if !isInGrid(grid, x, y) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	total := 1
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		total += areaOfIsland(grid, nx, ny)
	}
	return total
}
