package main

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	length := len(matrix) * len(matrix[0])
	top := 0
	left := 0
	right := len(matrix[0]) - 1
	down := len(matrix) - 1
	for len(ans) < length {
		for i := left; i <= right && len(ans) < length; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		for i := top; i <= down && len(ans) < length; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		for i := right; i >= left && len(ans) < length; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		for i := down; i >= top && len(ans) < length; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
	}
	return ans
}
