package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{8, 6, 2, 3, 1, 5, 4, 7, 9, 1}
	result := quickSort(arr)
	fmt.Println(result)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	// pivot := rand.Int() % len(arr)
	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] > arr[right] {
			fmt.Println(i, left)
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr

}

func intervalsQuickSort(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	left, right := 0, len(intervals)-1

	// pivot := rand.Int() % len(intervals)
	pivot := rand.Int() % len(intervals)
	intervals[right], intervals[pivot] = intervals[pivot], intervals[right]

	for i := range intervals {
		if intervals[i][1] < intervals[right][1] {
			intervals[i], intervals[left] = intervals[left], intervals[i]
			left++
		}
	}
	intervals[left], intervals[right] = intervals[right], intervals[left]
	intervalsQuickSort(intervals[:left])
	intervalsQuickSort(intervals[left+1:])
	return intervals

}

func eraseOverlapIntervals(intervals [][]int) int {
	intervals = intervalsQuickSort(intervals)
	count := 1
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
	}
	return len(intervals) - count
}
