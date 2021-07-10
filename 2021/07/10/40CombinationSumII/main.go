package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	dfs(&result, target, 0, []int{}, candidates)
	return result
}

func dfs(result *[][]int, target int, start int, current []int, candidates []int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
	}

	for i := start; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if target-candidates[i] >= 0 {
			current = append(current, candidates[i])
			dfs(result, target-candidates[i], i+1, current, candidates)
			current = current[:len(current)-1]
		}
	}

}
