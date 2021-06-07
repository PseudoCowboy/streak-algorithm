package main

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(dp)-1]
}

func minCostClimbingStairs1(cost []int) int {
	var cur, last int
	for i := 2; i < len(cost)+1; i++ {
		cur, last = last, min(cur+cost[i-2], last+cost[i-1])
	}
	return last
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
