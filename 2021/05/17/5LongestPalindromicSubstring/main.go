package main

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, left, right int, res string) string {
	sub := ""
	for left >= 0 && right < len(s) && s[left] == s[right] {
		sub = s[left : right+1]
		left--
		right++
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}

func longestPalindrome2(s string) string {
	res := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i) < 3 || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}

	return res
}
