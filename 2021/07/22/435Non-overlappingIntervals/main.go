package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	cur := intervals[len(intervals)-1][0]
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i][1] > cur {
			count++
		} else {
			cur = intervals[i][0]
		}
	}

	return count
}
