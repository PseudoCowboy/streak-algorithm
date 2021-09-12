package main

import "math/rand"

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
				continue
			}
			if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
				continue
			}
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	ans := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if triangle[len(triangle)-1][i] < ans {
			ans = triangle[len(triangle)-1][i]
		}
	}
	rand.Intn(7)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
