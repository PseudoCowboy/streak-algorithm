package main

import "fmt"

func main() {
	input := 123

	fmt.Println(reverse(input))
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		remainder := x % 10
		x /= 10
		result = remainder + result*10
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}
