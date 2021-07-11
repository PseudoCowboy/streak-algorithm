package main

import "fmt"

func main() {
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
}

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	dfs(&result, []int{}, 0, nums)
	return result
}

func dfs(result *[][]int, current []int, start int, nums []int) {
	if len(current) >= 2 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
	}
	used := make([]bool, 201)
	for i := start; i < len(nums); i++ {
		if (len(current) == 0 || nums[i] >= current[len(current)-1]) && !used[nums[i]+100] {
			used[nums[i]+100] = true
			current = append(current, nums[i])
			dfs(result, current, i+1, nums)
			current = current[:len(current)-1]
		}
	}
}
