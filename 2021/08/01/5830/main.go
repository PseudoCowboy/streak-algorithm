package main

import "fmt"

func main() {
	fmt.Println(isThree(2))
}

func isThree(n int) bool {
	if n == 1 {
		return false
	}
	count := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 1
}
