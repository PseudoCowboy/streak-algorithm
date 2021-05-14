package main

func jump(nums []int) int {
	length := len(nums)
	memo := make([]int, length)
	for i := range memo {
		memo[i] = length
	}
	return dp(nums, &memo, 0)
}

func dp(nums []int, memo *[]int, p int) int {
	if p >= len(nums)-1 {
		return 0
	}

	if (*memo)[p] != len(nums) {
		return (*memo)[p]
	}
	steps := nums[p]
	for i := 1; i <= steps; i++ {
		subProblem := dp(nums, memo, p+i)
		(*memo)[p] = min((*memo)[p], subProblem+1)
	}
	return (*memo)[p]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
