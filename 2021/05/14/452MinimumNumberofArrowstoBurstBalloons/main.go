package main

import "math/rand"

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

func findMinArrowShots(points [][]int) int {
	points = intervalsQuickSort(points)
	count := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			count++
			end = points[i][1]
		}
	}
	return count
}
