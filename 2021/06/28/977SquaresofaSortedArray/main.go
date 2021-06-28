package main

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	low, high, total := 0, len(nums)-1, len(nums)-1
	for low <= high {
		h := nums[high] * nums[high]
		l := nums[low] * nums[low]
		if h > l {
			result[total] = h
			high--
		} else {
			result[total] = l
			low++
		}
		total--
	}

	return result
}
