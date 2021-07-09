package main

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, n)
	dfs(n, k, 0, []int{}, &used, &result)
	return result
}

func dfs(n, k, start int, current []int, used *[]bool, result *[][]int) {
	if k == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
	}

	for i := start; i < 9; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			current = append(current, i+1)
			dfs(n, k-1, i+1, current, used, result)
			(*used)[i] = false
			current = current[:len(current)-1]
		}
	}
}
