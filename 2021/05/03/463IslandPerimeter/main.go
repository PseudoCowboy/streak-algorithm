package main

func islandPerimeter(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i-1 < 0 || grid[i-1][j] == 0 {
					area++
				}
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					area++
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					area++
				}
				if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
					area++
				}
			}
		}
	}
	return area
}
