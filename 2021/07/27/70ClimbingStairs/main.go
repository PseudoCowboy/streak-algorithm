package main

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	first, second := 1, 1
	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return first + second
}
