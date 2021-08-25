package main

import "fmt"

func main() {
	fmt.Println(judgeSquareSum(10))
}

func judgeSquareSum(c int) bool {
	if c == 0 {
		return true
	}
	left, right := 0, sqrt(c)
	fmt.Println(right)
	for left <= right {
		leftSquare := left * left
		rightSquare := right * right
		if leftSquare+rightSquare == c {
			return true
		} else if leftSquare+rightSquare < c {
			left++
		} else {
			right--
		}
	}

	return false
}

func sqrt(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		fmt.Println(left, right, mid)
		if n/mid == mid && n%mid == 0 {
			return mid
		}
		if n/mid == mid && n%mid > 0 {
			left = mid + 1
		}
		if n/mid < mid {
			right = mid - 1
		}
		if n/mid > mid {
			left = mid + 1
		}
	}
	return right
}
