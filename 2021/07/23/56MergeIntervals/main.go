package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= pre[1] {
			if intervals[i][1] > pre[1] {
				pre[1] = intervals[i][1]
			}
		} else {
			ans = append(ans, pre)
			pre = intervals[i]
		}
	}

	ans = append(ans, pre)
	return ans
}
