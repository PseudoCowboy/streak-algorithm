package main

import "fmt"

func main() {
	minSpaceWastedKResizing([]int{10, 20, 15, 30, 20}, 2)
}

func minSpaceWastedKResizing(nums []int, k int) int {
	max := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	diff := make([]int, len(nums))
	for i := range nums {
		diff[i] = max - nums[i]
	}
	fmt.Println(diff)
	return 0
}
