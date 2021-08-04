package main

func numSquares(n int) int {
	squares := make([]int, 0)
	for i := 1; i < n; i++ {
		square := i * i
		if square > n {
			break
		}
		squares = append(squares, square)
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
	}
	dp[0] = 0
	for i := range squares {
		for j := 1; j <= n; j++ {
			if j >= squares[i] {
				dp[j] = min(dp[j], dp[j-squares[i]]+1)
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
