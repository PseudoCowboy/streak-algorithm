package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	dfs([]int{}, nums, &result, 0)
	return result
}

func dfs(current []int, nums []int, result *[][]int, start int) {
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		dfs(current, nums, result, i+1)
		current = current[:len(current)-1]
	}
}
