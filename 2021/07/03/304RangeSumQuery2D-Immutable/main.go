package main

type NumMatrix struct {
	matrixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	matrixSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sub := make([]int, len(matrix[0]))
		sub[0] = matrix[i][0]
		for j := 1; j < len(matrix[0]); j++ {
			sub[j] = sub[j-1] + matrix[i][j]
		}
		matrixSum[i] = sub
	}
	return NumMatrix{matrixSum: matrixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			sum += this.matrixSum[i][col2]
		} else {
			sum += this.matrixSum[i][col2] - this.matrixSum[i][col1-1]
		}

	}

	return sum
}

type NumMatrix1 struct {
	matrixSum [][]int
}

func Constructor1(matrix [][]int) NumMatrix1 {
	matrixSum := make([][]int, len(matrix)+1)
	matrixSum[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		matrixSum[i+1] = make([]int, len(matrix[0])+1)
		for j := range matrix[0] {
			matrixSum[i+1][j+1] = matrixSum[i][j+1] + matrixSum[i+1][j] - matrixSum[i][j] + matrix[i][j]
		}
	}
	return NumMatrix1{matrixSum: matrixSum}
}

func (this *NumMatrix1) SumRegion1(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrixSum[row2+1][col2+1] - this.matrixSum[row2+1][col1] - this.matrixSum[row1][col2+1] + this.matrixSum[row1][col1]
}
