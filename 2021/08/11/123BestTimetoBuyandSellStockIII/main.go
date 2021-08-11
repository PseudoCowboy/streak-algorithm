package main

import "fmt"

func main() {
	maxProfit([]int{2, 4, 1})
}

func maxProfit(prices []int) int {
	skip := make([]int, 2)
	sum := 0
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			skip[0] = i
			min = prices[i]
		}
		if prices[i]-min > max {
			skip[1] = i
			max = prices[i] - min
		}
	}
	fmt.Println(max, skip)
	sum += max
	if max == 0 || len(prices) < 3 {
		return sum
	}
	max = 0
	if skip[0] == 0 && skip[1] == 1 {
		min = prices[2]
	} else if skip[0] == 0 {
		min = prices[1]
	} else {
		min = prices[0]
	}
	for i := 1; i < len(prices); i++ {
		if i >= skip[0] && i <= skip[1] {
			continue
		}
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return sum + max
}
