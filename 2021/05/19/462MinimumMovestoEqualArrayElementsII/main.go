package main

import "math/rand"

func minMoves2(nums []int) int {
	nums = quickSort(nums)
	midVal := nums[len(nums)/2]
	count := 0
	for i := 0; i < len(nums); i++ {
		delta := nums[i] - midVal
		if delta < 0 {
			count -= delta
		} else {
			count += delta
		}
	}

	return count
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1
	pivot := rand.Int() % len(nums)

	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	quickSort(nums[:left])
	quickSort(nums[left+1:])
	return nums
}
