package main

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	result := make([][]string, 0)
	used := make([]bool, n)
	level := 0
	helper()
	return result
}

func helper(n, level int, result *[][]string, current []string) {
	if level == n {
		temp := make([]string, n)
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < n; i++ {
		row := generate(n)
		row[i]
		helper(n, level+1, result, current)
	}
}

func generate(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "."
	}
	return str
}
