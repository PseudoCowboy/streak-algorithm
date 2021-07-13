package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
