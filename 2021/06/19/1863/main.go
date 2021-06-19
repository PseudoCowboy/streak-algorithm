package main

func subsetXORSum(nums []int) int {
	sum := 0
	dfs(nums, 0, 0, &sum)
	return sum
}

func dfs(nums []int, index, current int, sum *int) {
	if index == len(nums) {
		*sum += current
		return
	}

	dfs(nums, index+1, current, sum)
	dfs(nums, index+1, current^nums[index], sum)
}
