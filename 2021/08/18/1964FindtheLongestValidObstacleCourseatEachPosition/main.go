package main

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	dp := make([]int, 0)
	for i, obstacle := range obstacles {
		if len(dp) == 0 || dp[len(dp)-1] <= obstacle {
			dp = append(dp, obstacle)
			ans[i] = len(dp)
		} else {

		}
	}

	return ans
}
