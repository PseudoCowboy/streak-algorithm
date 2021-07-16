package main

func maxProfit(prices []int) int {
	sum := 0
	pre := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < pre {
			pre = prices[i]
		} else {
			sum += prices[i] - pre
			pre = prices[i]
		}
	}
	return sum
}
