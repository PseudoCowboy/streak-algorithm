package main

import (
	"sort"
)

func main() {
	deleteAndEarn([]int{10, 8, 4, 2, 1, 3, 4, 8, 2, 9, 10, 4, 8, 5, 9, 1, 5, 1, 6, 8, 1, 1, 6, 7, 8, 9, 1, 7, 6, 8, 4, 5, 4, 1, 5, 9, 8, 6, 10, 6, 4, 3, 8, 4, 10, 8, 8, 10, 6, 4, 4, 4, 9, 6, 9, 10, 7, 1, 5, 3, 4, 4, 8, 1, 1, 2, 1, 4, 1, 1, 4, 9, 4, 7, 1, 5, 1, 10, 3, 5, 10, 3, 10, 2, 1, 10, 4, 1, 1, 4, 1, 2, 10, 9, 7, 10, 1, 2, 7, 5})
}

func deleteAndEarn(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] += nums[i]
	}
	if len(m) == 1 {
		return m[nums[0]]
	}
	arr := make([]int, 0)
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	dp := make([]int, len(arr))
	dp[0] = m[arr[0]]
	if arr[0]+1 == arr[1] {
		dp[1] = max(m[arr[0]], m[arr[1]])
	} else {
		dp[1] = m[arr[0]] + m[arr[1]]
	}

	for i := 2; i < len(arr); i++ {
		if arr[i-1]+1 == arr[i] {
			dp[i] = max(dp[i-2]+m[arr[i]], dp[i-1])
		} else {
			dp[i] = dp[i-1] + m[arr[i]]
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
