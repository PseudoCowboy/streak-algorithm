package main

import "sort"

type schedule struct {
	s   int
	e   int
	pro int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	arr := make([]schedule, len(startTime))
	for i := range startTime {
		arr[i] = schedule{
			s:   startTime[i],
			e:   endTime[i],
			pro: profit[i],
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].e == arr[j].e {
			return arr[i].pro < arr[j].pro
		}
		return arr[i].e < arr[j].e
	})
	dp := make([]int, len(arr))
	dp[0] = arr[0].pro
	for i := 1; i < len(arr); i++ {
		left, right := 0, i-1
		for left < right {
			mid := left + (right-left)>>1
			if arr[mid+1].e <= arr[right].s {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if arr[left].e <= arr[i].s {
			dp[i] = max(dp[i-1], dp[left]+arr[i].pro)
		} else {
			dp[i] = max(dp[i-1], arr[i].pro)
		}
	}
	return dp[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
