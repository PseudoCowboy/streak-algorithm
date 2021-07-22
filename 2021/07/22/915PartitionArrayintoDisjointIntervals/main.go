package main

import "fmt"

func main() {

	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
}

func partitionDisjoint(nums []int) int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = nums[0]
	right[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		left[i] = max(left[i-1], nums[i])
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = min(right[i+1], nums[i])
	}
	for i := range left {
		if left[i] <= right[i+1] {
			return i + 1
		}
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
