package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, &result, []int{}, &used)
	return result
}

func dfs(nums []int, result *[][]int, current []int, used *[]bool) {
	if len(current) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			current = append(current, nums[i])
			dfs(nums, result, current, used)
			(*used)[i] = false
			current = current[:len(current)-1]

		}
	}
}
