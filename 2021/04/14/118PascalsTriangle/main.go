package main

import "fmt"

func main() {
	fmt.Print(generate(3))

}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 1 {
		result[0] = []int{1}
		return result
	}
	if numRows == 2 {
		result[0] = []int{1}
		result[1] = []int{1, 1}
		return result
	}
	result[0] = []int{1}
	result[1] = []int{1, 1}
	for i := range result {
		if i == 0 || i == 1 {
			continue
		}
		result[i] = triangle(result[i-1], i+1)
	}
	return result
}

func triangle(pre []int, length int) []int {
	result := make([]int, length)
	for i := range result {
		if i == 0 || i == length-1 {
			result[i] = 1
			continue
		}
		result[i] = pre[i-1] + pre[i]
	}
	return result
}
