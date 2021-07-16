package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		if i > total {
			return false
		}
		total = max(total, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
