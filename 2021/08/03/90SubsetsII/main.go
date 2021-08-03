package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	dfs(&ans, []int{}, nums, 0)
	return ans
}

func dfs(ans *[][]int, current []int, nums []int, index int) {
	temp := make([]int, len(current))
	copy(temp, current)
	*ans = append(*ans, temp)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		current = append(current, nums[i])
		dfs(ans, current, nums, i+1)
		current = current[:len(current)-1]
	}
}
