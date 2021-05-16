package main

func rob(nums []int) int {
	return dp(nums, 0)
}

func dp(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	return max(dp(nums, start+1), nums[start]+dp(nums, start+2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法一 DP
func rob198(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// dp[i] 代表抢 nums[0...i] 房子的最大价值
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// 解法二 DP 优化辅助空间，把迭代的值保存在 2 个变量中
func rob198_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	curMax, preMax := 0, 0
	for i := 0; i < n; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}
