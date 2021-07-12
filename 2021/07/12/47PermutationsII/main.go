package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	dfs(&result, []int{}, nums, &used)
	return result
}

func dfs(result *[][]int, current, nums []int, used *[]bool) {
	if len(current) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
		}
		if !(*used)[i] {
			(*used)[i] = true
			current = append(current, nums[i])
			dfs(result, current, nums, used)
			(*used)[i] = false
			current = current[:len(current)-1]
		}
	}
}
