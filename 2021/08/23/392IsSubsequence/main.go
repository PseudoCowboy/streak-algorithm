package main

func isSubsequence1(s string, t string) bool {
	sindex := 0
	tindex := 0
	for tindex < len(t) {
		if t[tindex] == s[sindex] {
			sindex++
			if sindex == len(s) {
				return true
			}
		}
		tindex++
	}
	return sindex == len(s)
}

func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(t)][len(s)] == len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
