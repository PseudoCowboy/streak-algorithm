package main

import "fmt"

func main() {
	maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4})
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}

	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	}

	for i := 1; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
		dp[i][0] = max(dp[i-1][0], dp[i][0])
	}

	for i := 1; i < len(nums2); i++ {
		if nums2[i] == nums1[0] {
			dp[0][i] = 1
		}
		dp[0][i] = max(dp[0][i-1], dp[0][i])
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)

	return dp[len(dp)-1][len(dp[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
