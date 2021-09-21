package main

import "fmt"

func main() {
	fmt.Println(sumOfBeauties([]int{3, 2, 1}))
}

func sumOfBeauties(nums []int) int {
	sum := make([]int, len(nums))
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = max(prefix[i-1], nums[i])
	}
	postfix[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		postfix[i] = min(postfix[i+1], nums[i])
	}
	for i := 1; i < len(nums)-1; i++ {
		if prefix[i-1] < nums[i] && postfix[i+1] > nums[i] {
			sum[i] = 2
		} else if nums[i-1] < nums[i] && nums[i+1] > nums[i] {
			sum[i] = 1
		}
	}
	ans := 0
	for _, v := range sum {
		ans += v
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
