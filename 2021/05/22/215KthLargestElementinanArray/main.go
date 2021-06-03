package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[k-1]
}

func main() {
	arr := []int{7, 8, 9, 4, 3, 2, 1, 5, 6, 10, 11}
	sortK(arr, 3)
	fmt.Println(arr)
}

func sortK(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	rand.Seed(time.Now().Unix())
	pivot := rand.Int() % len(nums)
	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]
	fmt.Println(pivot, left, k, nums)
	if left == k-1 {
		return
	}
	if left < k-1 {
		sortK(nums[left+1:], k-1-left)
	}
	if left > k-1 {
		sortK(nums[:left], k)
	}
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
