package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findGCD([]int{2, 5, 6, 9, 10}))
}

func findGCD(nums []int) int {
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	ans := 1
	for i := min; i >= 1; i-- {
		if min%i == 0 && max%i == 0 {
			return i
		}
	}
	return ans
}
