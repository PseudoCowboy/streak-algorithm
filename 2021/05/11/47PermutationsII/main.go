package main

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	current := make([]int, 0)
	used := make([]bool, len(nums))
	level := 0
	sort.Ints(nums)
	helper(nums, level, current, &result, &used)
	return result
}

func helper(nums []int, level int, current []int, result *[][]int, used *[]bool) {
	if level == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	// 1, 1, 3
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			current = append(current, nums[i])
			(*used)[i] = true
			helper(nums, level+1, current, result, used)
			current = current[:len(current)-1]
			(*used)[i] = false
		}
	}
}
