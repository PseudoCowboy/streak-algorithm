package main

func maxProfit(k int, prices []int) int {
	dp := make([]int, 2*k)
	for i := 0; i < len(dp); i += 2 {
		dp[i] = -prices[0]
	}

	for _, price := range prices {
		for i := len(dp) - 1; i >= 0; i-- {
			if i%2 == 1 {
				if i == 1 {
					dp[i] = max(dp[i], price+dp[i-1])
					continue
				}
				dp[i] = max(dp[i], price+dp[i-1])
			} else {
				if i == 0 {
					dp[i] = max(dp[i], -price)
					continue
				}
				dp[i] = max(dp[i], dp[i-1]-price)
			}
		}
	}

	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
