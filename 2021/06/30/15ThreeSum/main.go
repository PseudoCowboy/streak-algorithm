package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		two := twoSum(nums[i+1:], 0-nums[i])
		if len(two) != 0 {
			for j := range two {
				ans = append(ans, append(two[j], nums[i]))
			}
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func twoSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			ans = append(ans, []int{nums[left], nums[right]})
		} else if nums[left]+nums[right] < target {
			for left+1 < right && nums[left] == nums[left+1] {
				left++
			}
			left++
			continue
		} else if nums[left]+nums[right] > target {
			for right-1 > left && nums[right-1] == nums[right] {
				right--
			}
			right--
			continue
		}
		for left+1 < right && nums[left] == nums[left+1] {
			left++
		}
		for right-1 > left && nums[right-1] == nums[right] {
			right--
		}
		left++
		right--
	}
	return ans
}
