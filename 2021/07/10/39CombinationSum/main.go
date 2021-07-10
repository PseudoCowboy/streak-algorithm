package main

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	dfs(&result, candidates, 0, target, []int{})
	return result
}

func dfs(result *[][]int, candidates []int, index, target int, current []int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			current = append(current, candidates[i])
			dfs(result, candidates, i, target-candidates[i], current)
			current = current[:len(current)-1]
		}
	}
}
