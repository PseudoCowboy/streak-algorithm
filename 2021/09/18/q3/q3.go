package main

import "fmt"

func main() {

	fmt.Println(maxTaxiEarnings(20, [][]int{
		{2, 5, 4}, {1, 5, 1},
	}))
}

type r struct {
	from int
	tip  int
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	m := make(map[int][]r)
	for _, v := range rides {
		m[v[1]] = append(m[v[1]], r{
			from: v[0],
			tip:  v[2],
		})
	}
	fmt.Println(m)
	dp := make([]int64, n+1)
	for i := 1; i < len(dp); i++ {
		if val, ok := m[i]; ok {
			for _, item := range val {
				dp[i] = max3(dp[item.from]+int64(i-item.from)+int64(item.tip), dp[i-1], dp[i])
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}

func max3(a, b, c int64) int64 {
	if a > b {
		return max(a, c)
	}
	return max(b, c)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
