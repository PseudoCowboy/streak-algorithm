package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = 1
		}
	}

	return false
}
