package main

func maximumUniqueSubarray(nums []int) int {
	m := make(map[int]int)
	left, right := 0, 0
	max, sum := 0, 0
	for right < len(nums) {
		val := m[nums[right]]
		if val == 0 {
			m[nums[right]]++
			sum += nums[right]
			if sum > max {
				max = sum
			}
			right++
		} else {
			for left < right {
				if nums[left] == nums[right] {
					m[nums[left]]--
					sum -= nums[left]
					left++
					break
				}
				sum -= nums[left]
				m[nums[left]]--
				left++
			}
		}
	}
	return max
}
