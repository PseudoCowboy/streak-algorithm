package main

import "fmt"

func main() {
	monotoneIncreasingDigits(10)
}

func monotoneIncreasingDigits(n int) int {
	arr := make([]int, 0)
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	fmt.Println(arr)
	return 0
}
