package main

import "fmt"

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))

}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	dp := make([]int, 0)
	for i, obstacle := range obstacles {
		if len(dp) == 0 || dp[len(dp)-1] <= obstacle {
			dp = append(dp, obstacle)
			ans[i] = len(dp)
		} else {
			left, right := 0, len(dp)-1
			pos := right
			for left <= right {
				mid := left + (right-left)>>1
				if dp[mid] > obstacle {
					pos = mid
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			dp[pos] = obstacle
			ans[i] = pos + 1
		}
	}

	return ans
}
