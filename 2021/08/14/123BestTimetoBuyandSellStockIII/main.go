package main

import "math"

func main() {
	maxProfit([]int{2, 4, 1})
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 5)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], prices[i]+dp[i-1][1])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit1(prices []int) int {
	firstbuy, firstsell := math.MinInt32, 0
	secondbuy, secondsell := math.MinInt32, 0
	for _, price := range prices {
		secondsell = max(secondsell, price+secondbuy)
		secondbuy = max(secondbuy, firstsell-price)
		firstsell = max(firstsell, price+firstbuy)
		firstbuy = max(firstbuy, -price)
	}
	return secondsell
}
