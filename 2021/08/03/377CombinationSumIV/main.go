package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	sort.Ints(nums)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := range nums {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}
