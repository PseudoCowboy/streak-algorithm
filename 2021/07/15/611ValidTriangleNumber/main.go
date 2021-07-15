package main

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	count := 0
	for i := len(nums) - 1; i >= 2; i-- {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
