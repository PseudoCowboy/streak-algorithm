package main

import (
	"fmt"
	"math/rand"
)

func isCovered(ranges [][]int, left int, right int) bool {
	intervalsQuickSort(ranges)
	max := 0
	for i := range ranges {
		if ranges[i][0] <= left && ranges[i][1] >= left && max == 0 {
			max = ranges[i][1]
		}
		if ranges[i][0] <= max+1 && ranges[i][1] >= max {
			max = ranges[i][1]
		}
	}

	fmt.Println(max)

	return max >= right

}

func intervalsQuickSort(intervals [][]int) {
	if len(intervals) < 2 {
		return
	}
	left, right := 0, len(intervals)-1

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
}
