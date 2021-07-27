package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first, second := 0, 1
	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return first + second
}
