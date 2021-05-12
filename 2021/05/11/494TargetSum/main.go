package main

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	level := 0
	dfs(nums, &result, target, level)
	return result
}

func dfs(nums []int, result *int, rest int, level int) {
	if level == len(nums) {
		if rest == 0 {
			*result++
		}
		return
	}
	rest += nums[level]
	dfs(nums, result, rest, level+1)
	rest -= nums[level]
	rest -= nums[level]
	dfs(nums, result, rest, level+1)
	rest += nums[level]
}

func findTargetSumWays1(nums []int, target int) int {
	return 0
}
