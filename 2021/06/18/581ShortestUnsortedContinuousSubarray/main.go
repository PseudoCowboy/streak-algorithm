package main

import "math"

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	left, right := -1, -1
	minR := math.MaxInt32
	maxL := math.MinInt32
	isSort := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			isSort = true
		}
		if isSort {
			minR = min(minR, nums[i])
		}
	}
	isSort = false
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			isSort = true
		}
		if isSort {
			maxL = max(maxL, nums[i])
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > minR {
			left = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < maxL {
			right = i
			break
		}
	}
	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
