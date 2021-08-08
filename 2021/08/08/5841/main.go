package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
}

func lengthOfLIS(nums []int) {
	tail := make([]int, len(nums))
	tail[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left, right := 0, len(nums)-1
		for left < right {
			mid := left + (right-left)/2
			if tail[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		tail[left] = nums[i]
	}
	fmt.Println(tail)

}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {

	dp := make([]int, 0)
	tail := make([]int, len(obstacles))
	res := 0
	for _, v := range obstacles {
		i := 0
		j := res
		for i < j {
			m := (i + j) / 2
			if tail[m] < v {
				i = m + 1
			} else {
				j = m
			}
			tail[i] = v
			if res == j {
				res++
			}
		}
		fmt.Println(tail)
		dp = append(dp, res)
	}

	return dp
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
