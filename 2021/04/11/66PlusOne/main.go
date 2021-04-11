package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}

	a = append([]int{0}, a...)
	fmt.Println(a)
}

func plusOne(digits []int) []int {
	extendFlag := false
	index := 0

	for len(digits)-1-index >= 0 {
		if digits[len(digits)-1-index]+1 > 9 {
			digits[len(digits)-1-index] = 0
		} else {
			digits[len(digits)-1-index] = digits[len(digits)-1-index] + 1
			break
		}

		if len(digits)-1-index == 0 {
			extendFlag = true
		}
		index++

	}

	if extendFlag == true {
		digits = append([]int{1}, digits...)
	}

	return digits
}
