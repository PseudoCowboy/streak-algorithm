package main

import (
	"sort"
)

func main() {
	minimumDifference([]int{87063, 61094, 44530, 21297, 95857, 93551, 9918}, 6)
}

func minimumDifference(nums []int, k int) int {
	if k == 1 || len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	if k == len(nums) {
		return nums[k-1] - nums[0]
	}
	min := make([]int, 0)
	for i := 0; i < len(nums)-k+1; i++ {
		min = append(min, nums[i+k-1]-nums[i])
	}

	sort.Ints(min)
	return min[0]
}
