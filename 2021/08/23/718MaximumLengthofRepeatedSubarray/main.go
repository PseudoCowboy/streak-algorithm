package main

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}
	max := 0
	for i := range dp {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
			max = 1
		}
	}
	for i := range dp[0] {
		if nums2[i] == nums1[0] {
			dp[0][i] = 1
			max = 1
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}
