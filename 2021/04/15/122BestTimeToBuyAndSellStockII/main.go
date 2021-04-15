package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	total := 0
	min := prices[0]

	for i := range prices {
		if prices[i] > min {
			total += prices[i] - min
			min = prices[i]
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	return total

}
