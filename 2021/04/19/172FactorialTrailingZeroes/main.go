package main

func main() {
	trailingZeroes(5)
}

func trailingZeroes(n int) int {
	if n/5 == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}
