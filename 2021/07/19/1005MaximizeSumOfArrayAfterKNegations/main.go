package main

import (
	"math/rand"
)

func largestSumAfterKNegations(nums []int, k int) int {
	quick(nums)
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k > 0 && k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

func quick(nums []int) {
	left, right := 0, len(nums)-1
	pivot := rand.Int() % len(nums)
	nums[right], nums[pivot] = nums[pivot], nums[right]
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) > abs(nums[right]) {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	quick(nums[:left])
	quick(nums[left+1:])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
