package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minOperations([][]int{
		{1, 2},
		{3, 4},
	}, 1))
}

func minOperations(grid [][]int, x int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 0
	}
	arr := make([]int, len(grid)*len(grid[0]))
	index := 0
	sum := 0
	for i := range grid {
		for j := range grid[i] {
			sum += grid[i][j]
			arr[index] = grid[i][j]
			index++
		}
	}
	avg := sum / len(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	t1 := (arr[0] - avg) / x
	avg = arr[0] - t1*x
	ans := 0
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i]-arr[j])%x != 0 {
				return -1
			}
		}
	}
	for i := range arr {
		ans += (arr[i] - avg) / x
	}
	return ans
}
