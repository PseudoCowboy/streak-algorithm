package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	sum := 0
	m := make(map[int]int)
	for sum != 1 {
		sum = 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = 1
		n = sum
	}
	return true
}
