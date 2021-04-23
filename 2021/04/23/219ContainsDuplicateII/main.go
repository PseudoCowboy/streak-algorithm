package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i := range nums {
		if val, ok := m[nums[i]]; ok {
			if i-val <= k {
				return true
			}
		}
		m[nums[i]] = i
	}

	return false
}
