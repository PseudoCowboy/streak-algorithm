package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(gridGame([][]int{
		{20, 3, 20, 17, 2, 12, 15, 17, 4, 15},
		{20, 10, 13, 14, 15, 5, 2, 3, 14, 3},
	}))
}

func gridGame(grid [][]int) int64 {
	prefix := make([]int, len(grid[0]))
	postfix := make([]int, len(grid[0]))
	prefix[0] = grid[1][0]
	for i := 1; i < len(grid[0]); i++ {
		prefix[i] = prefix[i-1] + grid[1][i]
	}
	postfix[len(grid[0])-1] = grid[0][len(grid[0])-1]
	for i := len(grid[0]) - 2; i >= 0; i-- {
		postfix[i] = grid[0][i] + postfix[i+1]
	}
	dp := make([]int64, len(grid[0]))
	dp[0] = int64(postfix[1])
	dp[len(grid[0])-1] = int64(prefix[len(grid[0])-2])
	for i := 1; i < len(grid[0])-1; i++ {
		dp[i] = max(int64(prefix[i-1]), int64(postfix[i+1]))
	}

	ans := int64(0)
	ans = math.MaxInt64
	for i := range dp {
		if dp[i] < ans {
			ans = dp[i]
		}
	}
	return ans
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
