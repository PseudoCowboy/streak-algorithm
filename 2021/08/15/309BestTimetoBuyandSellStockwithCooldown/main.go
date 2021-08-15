package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)
	dp := make([]int, n+1)
	minp := prices[0]
	for i := 1; i < len(dp); i++ {
		if i >= 2 {
			minp = min(minp, prices[i-1]-dp[i-2])
		}
		dp[i] = max(dp[i-1], prices[i-1]-minp)
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
