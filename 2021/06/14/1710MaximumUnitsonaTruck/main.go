package main

import "math/rand"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	res := 0
	sort(boxTypes)
	for i := 0; i < len(boxTypes); i++ {
		if truckSize >= boxTypes[i][0] {
			res += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			res += boxTypes[i][1] * truckSize
			break
		}
	}

	return res
}

func sort(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}

	pivot := rand.Int() % len(matrix)

	left, right := 0, len(matrix)-1
	matrix[pivot], matrix[right] = matrix[right], matrix[pivot]
	for i := 0; i < len(matrix); i++ {
		if matrix[i][1] > matrix[right][1] {
			matrix[i], matrix[left] = matrix[left], matrix[i]
			left++
		}
	}
	matrix[left], matrix[right] = matrix[right], matrix[left]
	sort(matrix[:left])
	sort(matrix[left+1:])
}
