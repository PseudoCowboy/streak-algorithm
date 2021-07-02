package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(mySqrt(i))
	}
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if x/mid == mid {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else if mid < x/mid {
			left = mid + 1
		}
	}
	return right
}

func mySqrt1(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	return right
}
