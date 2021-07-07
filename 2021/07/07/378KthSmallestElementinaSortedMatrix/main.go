package main

func kthSmallest(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	left := matrix[0][0]
	right := matrix[m-1][n-1]
	for left <= right {
		mid := left + (right-left)/2
		count := 0
		for i := 0; i < m; i++ {
			j := n - 1
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += j + 1
		}
		if count < k {
			left = mid + 1
		} else if count > k {
			right = mid - 1
		} else {
			return mid
		}

	}
	return left
}
