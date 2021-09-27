package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{0, 1, 2}))
}

func sortArrayByParity(nums []int) []int {
	if len(nums) == 2 {
		if nums[0]%2 == 1 && nums[1]%2 == 0 {
			return []int{nums[1], nums[0]}
		} else {
			return nums
		}
	}
	left, right := 0, len(nums)-1
	for left < right {
		for nums[left]%2 == 0 {
			left++
			if left == len(nums) || left > right {
				return nums
			}
		}
		for nums[right]%2 == 1 {
			right--
			if right < 0 || left > right {
				return nums
			}
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
