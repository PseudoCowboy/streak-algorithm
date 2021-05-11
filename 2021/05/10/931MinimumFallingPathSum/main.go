package main

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := range dp[0] {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 {
				dp[i][j] = min2(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
				continue
			}
			if j == len(matrix[0])-1 {
				dp[i][j] = min2(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
				continue
			}
			dp[i][j] = min3(dp[i-1][j], dp[i-1][j+1], dp[i-1][j-1]) + matrix[i][j]
		}
	}
	min := 1000

	for i := range dp[len(dp)-1] {
		if min > dp[len(dp)-1][i] {
			min = dp[len(dp)-1][i]
		}
	}
	return min
}

func min3(a, b, c int) int {
	if a < b {
		return min2(a, c)
	}

	return min2(b, c)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
