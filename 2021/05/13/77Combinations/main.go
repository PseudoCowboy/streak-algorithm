package main

func combine(n int, k int) [][]int {
	used := make([]bool, n)
	result := make([][]int, 0)
	dfs(used, []int{}, &result, k, 1, 0)
	return result
}

func dfs(used []bool, current []int, result *[][]int, k, start, index int) {
	if index == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i <= len(used); i++ {
		if used[i-1] {
			continue
		}
		used[i-1] = true
		current = append(current, i)
		dfs(used, current, result, k, i, index+1)
		used[i-1] = false
		current = current[:len(current)-1]
	}

}
