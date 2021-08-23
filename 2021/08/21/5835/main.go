package main

import (
	"fmt"
)

// {-9, 9, 9, -9},
// 		{-10, 10, 7, 5},
// 		{4, -8, -9, 6},
// 		{10, 2, -6, 1},
func main() {
	fmt.Println(maxMatrixSum1([][]int{
		{1, 2, 3},
		{-1, -2, -3},
		{1, 2, 3},
	}))
}

func maxMatrixSum1(matrix [][]int) int64 {
	count := 0
	min := 1000000
	sum := int64(0)
	current := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				count++
			}
			current = abs(matrix[i][j])
			if current < min {
				min = current
			}
			sum += int64(current)
		}
	}
	if count%2 == 1 {
		return sum - int64(min) - int64(min)
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
