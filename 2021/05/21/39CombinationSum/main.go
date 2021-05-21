package main

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	dfs(candidates, target, 0, []int{}, &result)
	return result
}

func dfs(candidates []int, target int, start int, current []int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
	}

	for i := start; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			current = append(current, candidates[i])
			dfs(candidates, target-candidates[i], i, current, result)
			current = current[:len(current)-1]
		}
	}
}
