package main

import "fmt"

func main() {
	input := 121

	fmt.Println(isPalindrome(input))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	result := 0
	for temp != 0 {
		remainder := temp % 10
		temp /= 10
		result = remainder + result*10
	}

	return result == temp

}
