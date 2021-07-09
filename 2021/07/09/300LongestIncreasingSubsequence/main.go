package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for i := range dp {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
