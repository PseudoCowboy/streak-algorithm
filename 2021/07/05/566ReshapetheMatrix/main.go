package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	arr := make([]int, 0)

	for i := range mat {
		arr = append(arr, mat[i]...)
	}

	result := make([][]int, 0)
	for i := 0; i < r; i++ {
		result = append(result, arr[i*c:(i+1)*c])
	}

	return result
}

func matrixReshape1(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	col, row := 0, 0
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := range mat {
		for j := range mat[i] {
			result[row][col] = mat[i][j]
			col++
			if col == c {
				col = 0
				row++
			}
		}
	}

	return result
}
