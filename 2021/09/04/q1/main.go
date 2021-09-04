package main

import "fmt"

func main() {
	fmt.Println(findMiddleIndex([]int{1}))
}

func findMiddleIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] + nums[i]
	}
	pos := make([]int, len(nums))
	pos[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		pos[i] = pos[i+1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if pos[i+1] == 0 {
				return 0
			}
			continue
		}
		if i == len(nums)-1 {
			if pre[len(nums)-2] == 0 {
				return len(nums) - 1
			}
			continue
		}
		if pre[i-1] == pos[i+1] {
			return i
		}
	}

	return -1
}
