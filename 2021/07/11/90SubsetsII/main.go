package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	dfs(&result, []int{}, nums, 0)
	return result
}

func dfs(result *[][]int, current []int, nums []int, start int) {
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		current = append(current, nums[i])
		dfs(result, current, nums, i+1)
		current = current[:len(current)-1]
	}
}
