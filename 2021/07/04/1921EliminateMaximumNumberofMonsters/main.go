package main

func eliminateMaximum(dist []int, speed []int) int {
	ans := 0
	count := 0
	n := len(dist)

	remain := make([]int, n)
	for i := 0; i < n; i++ {
		remain[i] = dist[i]
	}

	for count <= ans && count < n && ans < n {
		for j := 0; j < n; j++ {
			remain[j] -= speed[j]
			if remain[j] <= 0 {
				remain[j] = 1
				speed[j] = 0
				count++
			}
		}
		count--
		ans++
		if count > 0 {
			break
		}
	}
	return ans
}
