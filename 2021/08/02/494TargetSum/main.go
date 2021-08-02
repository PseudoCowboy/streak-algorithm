package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// check do nums sum can reach target
	if sum < target {
		return 0
	}
	// check dp result,  x - (sum - x) = target, x = (sum + target) >> 1
	if (sum+target)%2 == 1 {
		return 0
	}
	bagSize := (sum + target) >> 1
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for _, v := range nums {
		for i := bagSize; i >= v; i-- {
			dp[i] += dp[i-v]
		}
	}
	return dp[bagSize]
}

func findTargetSumWays1(nums []int, target int) int {
	result := 0
	level := 0
	sums := make([]int, len(nums))
	sums[len(nums)-1] = nums[len(nums)-1]
	// record sums for optimize
	for i := len(nums) - 2; i >= 0; i-- {
		sums[i] = nums[i] + sums[i+1]
	}
	dfs(nums, target, &result, level, sums)
	return result
}

func dfs(nums []int, target int, result *int, level int, sums []int) {
	if level == len(nums) {
		if target == 0 {
			*result++
		}
		return
	}
	// optimize if current value is bigger than left sum then return
	if abs(target) > sums[level] {
		return
	}

	dfs(nums, target+nums[level], result, level+1, sums)
	dfs(nums, target-nums[level], result, level+1, sums)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
