package main

func maxProfit(prices []int) int {
	sum := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			sum += prices[i] - min
		}
		min = prices[i]

	}
	return sum
}

func maxProfit1(prices []int) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}
	return dp[(len(prices)-1)%2][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
