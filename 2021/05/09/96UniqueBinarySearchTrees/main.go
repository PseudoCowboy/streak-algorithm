package main

func numTrees(n int) int {
	sol := make([]int, n+1)
	sol[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			sol[i] += (sol[j-1] * sol[i-j])
		}
	}
	return sol[n]
}
