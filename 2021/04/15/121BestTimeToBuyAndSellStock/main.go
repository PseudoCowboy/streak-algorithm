package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	minPrice := prices[0]

	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
	}

	return max
}
