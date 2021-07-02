package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	matrix := make([][]int, 0)
	for k, v := range m {
		matrix = append(matrix, []int{k, v})
	}
	quickSort(matrix)
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, matrix[i][0])
	}
	return ans

}

func quickSort(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}

	left, right := 0, len(matrix)-1
	for i := 0; i < len(matrix); i++ {
		if matrix[i][1] > matrix[right][1] {
			matrix[left], matrix[i] = matrix[i], matrix[left]
			left++
		}
	}

	matrix[left], matrix[right] = matrix[right], matrix[left]

	quickSort(matrix[:left])
	quickSort(matrix[left+1:])
}
