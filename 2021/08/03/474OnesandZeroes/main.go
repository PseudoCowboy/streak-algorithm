package main

func main() {
	findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
}

func findMaxForm(strs []string, m int, n int) int {
	om := make(map[string]int)
	zm := make(map[string]int)

	for i := range strs {
		o := 0
		z := 0
		for _, v := range strs[i] {
			if v == '0' {
				z++
			}
			if v == '1' {
				o++
			}
		}
		om[strs[i]] = o
		zm[strs[i]] = z
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, v := range strs {
		one := om[v]
		zero := zm[v]
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				if i >= zero && j >= one {
					dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
				}
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
