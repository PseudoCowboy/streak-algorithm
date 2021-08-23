package main

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {

			} else {

			}
		}
	}
	return dp[len(s)][len(t)]
}
