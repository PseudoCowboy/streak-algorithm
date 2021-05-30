package main

import (
	"fmt"
	"strconv"
)

func main() {

	maxValue("-132", 3)
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	firstValue, secondValue, targetValue := 0, 0, 0
	for i := 0; i < len(firstWord); i++ {
		firstValue = firstValue*10 + int(firstWord[i]-'a')
	}
	for i := 0; i < len(secondWord); i++ {
		secondValue = secondValue*10 + int(secondWord[i]-'a')
	}
	for i := 0; i < len(targetWord); i++ {
		targetValue = targetValue*10 + int(targetWord[i]-'a')
		fmt.Println(targetWord[i]-'a', targetValue, int(targetWord[i]-'a'))
	}
	return firstValue+secondValue == targetValue
}

func maxValue(n string, x int) string {
	index := 0
	fmt.Println(n[0] == '-')
	if n[0] == '-' {
		for i := 1; i < len(n); i++ {
			fmt.Println(n[i] - '0')
			if int(n[i]-'0') > x {
				index = i
				break
			}
			if i == len(n)-1 {
				index = len(n)
			}
		}

	} else {
		for i := 0; i < len(n); i++ {
			if int(n[i]-'0') < x {
				index = i
				break
			}
			if i == len(n)-1 {
				index = len(n)
			}
		}
	}

	return n[:index] + strconv.Itoa(x) + n[index:]
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	max := 0
	for i := 0; i < len(nums); i++ {
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
