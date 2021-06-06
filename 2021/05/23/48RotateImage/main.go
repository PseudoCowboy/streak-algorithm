package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}

func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < len(mat); i++ {
		for j := i; j < len(mat); j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat)/2; j++ {
			mat[i][j], mat[i][len(mat)-1-j] = mat[i][len(mat)-1-j], mat[i][j]
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}

	return true
}
