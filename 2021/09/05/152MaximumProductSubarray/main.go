package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		max := nums[i]
		for j := i - 1; j >= 0; j-- {
			sum *= nums[j]
			if sum > max {
				max = sum
			}
		}
		if max > ans {
			ans = max
		}
	}
	return ans
}
