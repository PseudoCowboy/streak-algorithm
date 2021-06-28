package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n

	for num <= tar {
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[i][bottom] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			result[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++
	}
	return result
}
