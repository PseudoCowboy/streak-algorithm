package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	half := sum >> 1
	dp := make([]int, half+1)
	for _, v := range stones {
		for i := half; i >= v; i-- {
			dp[i] = max(dp[i], dp[i-v]+v)
		}
	}
	return sum - 2*dp[half]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
