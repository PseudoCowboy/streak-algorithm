package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	maxVal := 0
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] > maxVal {
				maxVal = dp[i][j]
			}
		}
	}
	return len(word1) - maxVal + len(word2) - maxVal

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
