package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	current := make([]int, 0)
	used := make([]bool, len(nums))
	level := 0
	helper(nums, level, &used, &result, current)
	return result
}

func helper(nums []int, level int, used *[]bool, result *[][]int, current []int) {
	if level == len(nums) {
		temp := make([]int, len(*used))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			current = append(current, nums[i])
			helper(nums, level+1, used, result, current)
			current = current[:len(current)-1]
			(*used)[i] = false
		}
	}
}
