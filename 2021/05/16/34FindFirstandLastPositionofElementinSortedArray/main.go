package main

func searchRange(nums []int, target int) []int {
	index := search(nums, target)
	left, right := index, index
	for nums[left] == target {
		left--
	}
	for nums[right] == target {
		right++
	}

	return []int{left, right}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
