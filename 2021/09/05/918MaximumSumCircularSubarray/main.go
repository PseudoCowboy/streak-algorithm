package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{5, -2, 5}))
}

func maxSubarraySumCircular(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 && nums[i] > sum {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	sum = nums[0]
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 && nums[i] < sum {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum < min {
			min = sum
		}
	}
	if max > total-min || total == min {
		return max
	}
	return total - min
}
