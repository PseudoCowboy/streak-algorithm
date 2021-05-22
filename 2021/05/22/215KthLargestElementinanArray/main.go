package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[k-1]
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := rand.Int() % len(nums)
	left, right := 0, len(nums)-1
	nums[right], nums[pivot] = nums[pivot], nums[right]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}
