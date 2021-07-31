package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	maxLeft[0] = height[0]
	maxRight[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	ans := 0

	for i := range height {
		if i == 0 || i == len(height)-1 {
			continue
		}
		diff := min(maxLeft[i-1], maxRight[i+1]) - height[i]
		if diff > 0 {
			ans += diff
		}
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

func trap1(height []int) int {
	ans := 0
	for i := range height {
		if i == 0 || i == len(height)-1 {
			continue
		}
		left := 0
		right := 0
		for j := i - 1; j >= 0; j-- {
			left = max(left, height[j])
		}
		for j := i + 1; j < len(height); j++ {
			right = max(right, height[j])
		}
		diff := min(left, right) - height[i]
		if diff > 0 {
			ans += diff
		}
	}
	return ans
}
