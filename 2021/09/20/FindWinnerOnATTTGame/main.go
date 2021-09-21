package main

func tictactoe(moves [][]int) string {
	if len(moves) < 5 {
		return "pending"
	}
	matrix := make([][]string, 3)
	for i := range matrix {
		matrix[i] = make([]string, 3)
	}
	for i := 0; i < len(moves); i++ {
		move := moves[i]
		if i%2 == 0 {
			matrix[move[0]][move[1]] = "A"
		} else {
			matrix[move[0]][move[1]] = "B"
		}
		if i >= 4 {
			ans := check(matrix)
			if ans != "" {
				return ans
			}
		}
	}
	return "draw"
}

func check(matrix [][]string) string {
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if matrix[i][j] != matrix[i][j-1] || matrix[i][j] == "" {
				break
			}
			if j == 2 {
				return matrix[i][j]
			}
		}
	}
	for i := 1; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] != matrix[i-1][j] || matrix[i][j] == "" {
				break
			}
			if i == 2 {
				return matrix[i][j]
			}
		}
	}
	if matrix[0][0] == matrix[1][1] && matrix[2][2] == matrix[0][0] && matrix[0][0] != "" {
		return matrix[0][0]
	}
	if matrix[2][0] == matrix[1][1] && matrix[1][1] == matrix[0][2] && matrix[1][1] != "" {
		return matrix[1][1]
	}
	return ""
}
