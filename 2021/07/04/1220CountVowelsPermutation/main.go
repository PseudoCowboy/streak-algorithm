package main

import "fmt"

func main() {
	fmt.Println(countVowelPermutation(144))
}

func countVowelPermutation(n int) int {
	v := []byte{'a', 'e', 'i', 'o', 'u'}
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, 5)
	}

	dp[1] = []int64{1, 1, 1, 1, 1}
	var mod int64 = 1000000007

	for i := 2; i <= n; i++ {
		for j := 0; j < 5; j++ {
			switch v[j] {
			case 'a':
				dp[i][j] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % mod
			case 'e':
				dp[i][j] = (dp[i-1][0] + dp[i-1][2]) % mod
			case 'i':
				dp[i][j] = (dp[i-1][1] + dp[i-1][3]) % mod
			case 'o':
				dp[i][j] = dp[i-1][2] % mod
			case 'u':
				dp[i][j] = (dp[i-1][2] + dp[i-1][3]) % mod
			}
		}
	}
	return int((dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]) % mod)
}
