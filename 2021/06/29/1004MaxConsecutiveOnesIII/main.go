package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	count := 0
	total := 0
	for right < len(nums) {
		total += nums[right]
		for right-left+1 > total+k {
			total -= nums[left]
			left++
		}
		count = max(count, right-left+1)
		right++
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
