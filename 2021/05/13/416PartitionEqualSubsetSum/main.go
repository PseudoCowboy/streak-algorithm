package main

import "fmt"

func main() {
	canPartition([]int{1, 5, 11, 5})
}

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	average := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, average+1)
	}
	for i := range dp {
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= average; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	fmt.Println(dp)
	return dp[len(nums)][average]

}

func canPartition1(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	average := sum / 2
	used := make([]bool, len(nums))

	return dfs(nums, used, 0, average, 2, 0)
}

func dfs(nums []int, used []bool, bucket int, average int, rest int, start int) bool {
	if rest == 0 {
		return true
	}
	if bucket == average {
		return dfs(nums, used, 0, average, rest-1, 0)
	}
	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if bucket+nums[i] > average {
			continue
		}
		used[i] = true
		bucket += nums[i]
		if dfs(nums, used, bucket, average, rest, i+1) {
			return true
		}
		used[i] = false
		bucket -= nums[i]
	}

	return false
}
