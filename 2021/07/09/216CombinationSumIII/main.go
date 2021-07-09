package main

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, 9)
	dfs(k, n, 0, &result, &used, []int{})
	return result
}

func dfs(k, n, index int, result *[][]int, used *[]bool, current []int) {
	if k == 0 && n == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	if k == 0 {
		return
	}

	for i := index; i < 9; i++ {
		if !(*used)[i] {
			if n-i-1 < 0 {
				continue
			}
			(*used)[i] = true
			current = append(current, i+1)
			dfs(k-1, n-i-1, i+1, result, used, current)
			(*used)[i] = false
			current = current[:len(current)-1]
		}
	}
}
