package main

func maximumDifference(nums []int) int {
	ans := -1
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else {
			if nums[i]-min > ans {
				ans = nums[i] - min
			}
		}
	}
	return ans
}
