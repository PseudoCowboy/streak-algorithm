package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	far := 0
	for i, v := range nums {
		if i > far {
			return false
		}
		far = max(far, i+v)
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
