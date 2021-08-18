package main

func main() {
	lengthOfLIS1([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	max := 0
	for i := range dp {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS1(nums []int) int {
	dp := make([]int, 0)
	for _, num := range nums {
		if len(dp) == 0 || num > dp[len(dp)-1] {
			dp = append(dp, num)
		} else {
			left, right := 0, len(dp)-1
			pos := right
			for left <= right {
				mid := left + (right-left)>>1
				if dp[mid] >= num {
					pos = mid
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			dp[pos] = num
		}
	}
	return len(dp)
}
