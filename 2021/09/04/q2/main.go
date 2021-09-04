package main

import "fmt"

func main() {
	fmt.Println(findFarmland([][]int{
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		// {0},
		// {0},
		// {1},
		// {1},
		// {1},
		// {0},
		// {1},
		// {1},
	}))
}

func findFarmland(land [][]int) [][]int {
	ans := make([][]int, 0)
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 {
				endi, endj := find(land, i, j)
				ans = append(ans, []int{i, j, endi, endj})
			}
		}
	}
	return ans
}

func find(land [][]int, i, j int) (x, y int) {
	starti, startj := i, j
	for i < len(land) {
		if land[i][startj] == 1 {
			i++
		} else {
			break
		}
	}

	for j < len(land[0]) {
		if land[starti][j] == 1 {
			j++
		} else {
			break
		}
	}

	for p := starti; p < i; p++ {
		for q := startj; q < j; q++ {
			land[p][q] = 0
		}
	}
	return i - 1, j - 1
}
