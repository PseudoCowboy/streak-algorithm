package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	q := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
		if i-k+1 >= q[0] {
			q = q[1:]
		}
	}
	return ans
}
