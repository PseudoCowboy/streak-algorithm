package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	last := -1
	temp := 0
	ans := 0
	for i := range nums {
		if nums[i] > right {
			last = i
		}
		if nums[i] >= left {
			temp = i - last
		}
		ans += temp
	}
	return ans
}
