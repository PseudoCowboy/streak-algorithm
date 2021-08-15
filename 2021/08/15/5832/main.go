package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeArray([]int{6, 2, 0, 9, 7}))
}

func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	target := make([]int, 0)
	flag := true
	for left <= right {
		if flag {
			target = append(target, nums[left])
			left++
			flag = false
		} else {
			target = append(target, nums[right])
			right--
			flag = true
		}
	}
	return target
}
