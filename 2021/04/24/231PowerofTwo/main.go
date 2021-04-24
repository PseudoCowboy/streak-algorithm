package main

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	val := 1
	for true {
		val *= 2
		if val == n {
			return true
		}
		if val > n {
			return false
		}
	}
	return false
}
