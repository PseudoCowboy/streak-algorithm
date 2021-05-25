package main

import (
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// 暴力递归超时
func numDecodings1(s string) int {
	result := 0
	helper(s, &result)
	return result
}

func helper(s string, result *int) {
	if len(s) == 0 {
		*result++
		return
	}
	if s[0] == '0' {
		return
	}
	if s[0] == '1' {
		helper(s[1:], result)
		if len(s) > 1 {
			helper(s[2:], result)
		}
	} else if s[0] == '2' {
		helper(s[1:], result)
		if len(s) > 1 && s[1] <= '6' {
			helper(s[2:], result)
		}
	} else {
		helper(s[1:], result)
	}

}
