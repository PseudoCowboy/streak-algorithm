package main

func minFallingPathSum(matrix [][]int) int {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j+1])
				continue
			}
			if j == len(matrix[0])-1 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j-1])
				continue
			}
			matrix[i][j] += min(min(matrix[i-1][j-1], matrix[i-1][j+1]), matrix[i-1][j])
		}
	}
	ans := matrix[len(matrix)-1][0]
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[len(matrix)-1][i] < ans {
			ans = matrix[len(matrix)-1][i]
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
