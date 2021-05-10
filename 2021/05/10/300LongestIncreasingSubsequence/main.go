package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
