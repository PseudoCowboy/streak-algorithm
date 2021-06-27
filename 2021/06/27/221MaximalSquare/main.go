package main

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	maxSide := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		if i == 0 {
			for j := range dp[i] {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
					maxSide = 1
				}
			}
		}
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; i < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

func min3(a, b, c int) int {
	if a < b {
		return min(a, c)
	}
	return min(b, c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
