package main

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}
	isPal := make([][]bool, n)
	for i := range isPal {
		isPal[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || isPal[i+1][j-1]) {
				isPal[i][j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		if isPal[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPal[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func minCut1(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i - 1
	}

	for i := 0; i < n; i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			dp[i+j+1] = min(dp[i+j+1], dp[i-j]+1)
		}
		for j := 1; i-j+1 >= 0 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			dp[i+j+1] = min(dp[i+j+1], 1+dp[i-j+1])
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
