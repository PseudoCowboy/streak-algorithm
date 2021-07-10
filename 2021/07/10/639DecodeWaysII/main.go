package main

func numDecodings(s string) int {
	MOD := 1000000007
	s = " " + s
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == '*' {
			dp[i] = dp[i-1] * 9
			if s[i-1] == '1' || s[i-1] == '*' {
				dp[i] += dp[i-2] * 9
			}
			if s[i-1] == '2' || s[i-1] == '*' {
				dp[i] += dp[i-2] * 6
			}
		} else {
			if s[i] != '0' {
				dp[i] = dp[i-1]
			}
			if s[i-1] == '1' || s[i-1] == '*' {
				dp[i] += dp[i-2]
			}
			if s[i] >= '0' && s[i] <= '6' && (s[i-1] == '2' || s[i-1] == '*') {
				dp[i] += dp[i-2]
			}
		}
		dp[i] %= MOD
	}

	return dp[len(s)-1] % MOD

}
