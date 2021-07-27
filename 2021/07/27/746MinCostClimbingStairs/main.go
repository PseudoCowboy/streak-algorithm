package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		if cost[0] < cost[1] {
			return cost[0]
		}
		return cost[1]
	}
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[len(cost)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
