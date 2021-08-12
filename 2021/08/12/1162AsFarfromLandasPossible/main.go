package main

func main() {
	maxDistance([][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	})
}

func maxDistance(grid [][]int) int {
	q := make([][]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	if len(q) == 0 || len(q) == len(grid)*len(grid) {
		return -1
	}

	dir := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	// fmt.Println(q)

	level := -1
	for len(q) != 0 {
		level++
		size := len(q)
		for i := range q {
			x := q[i][0]
			y := q[i][1]
			for j := range dir {
				nextx := x + dir[j][0]
				nexty := y + dir[j][1]
				if nextx >= 0 && nexty >= 0 && nextx < len(grid) && nexty < len(grid) && grid[nextx][nexty] == 0 {
					grid[nextx][nexty] = 1
					q = append(q, []int{nextx, nexty})
				}
			}
		}
		q = q[size:]
	}

	return level
}
