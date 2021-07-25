package main

func maxProfit(prices []int, fee int) int {
	ans := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i] > min+fee {
			ans += prices[i] - min - fee
			min = prices[i] - fee
		}
	}

	return ans
}
