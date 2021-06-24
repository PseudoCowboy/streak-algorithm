package main

func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	for i := 1; i < len(nums)+1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return len(nums) + 1
}
