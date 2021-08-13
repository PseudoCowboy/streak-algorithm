package main

func setZeroes(matrix [][]int) {
	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rowMap[i] = true
				colMap[j] = true
			}
		}
	}
	for k := range rowMap {
		for i := range matrix[k] {
			matrix[k][i] = 0
		}
	}

	for k := range colMap {
		for i := range matrix {
			matrix[i][k] = 0
		}
	}
}
