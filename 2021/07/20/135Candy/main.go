package main

func candy(ratings []int) int {
	ans := make([]int, len(ratings))
	ans[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			ans[i] = max(ans[i], ans[i+1]+1)
		}
	}

	sum := 0
	for i := range ans {
		sum += ans[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
