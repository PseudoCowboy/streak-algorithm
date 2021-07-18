package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{
		{1, 5},
		{3, 2},
		{4, 2},
	}))
}

func maxPoints(points [][]int) int64 {
	m := len(points)
	n := len(points[0])
	dp := make([]int64, n)
	ans := int64(0)
	for i := range dp {
		dp[i] = int64(points[0][i])
	}

	row := 1
	for row < m {
		for i := 0; i < n; i++ {
			temp := int64(0)
			for j := 0; j < n; j++ {
				temp = max(temp, int64(points[row][j])-abs(i, j))
			}
			dp[i] += temp
		}
		row++
	}

	for i := range dp {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int64 {
	if a-b > 0 {
		return int64(a - b)
	}
	return int64(b - a)
}
