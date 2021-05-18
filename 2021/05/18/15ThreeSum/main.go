package main

import (
	"math/rand"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	threeSum(arr)
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	nums = quickSort(nums)
	for start := 0; start < len(nums)-1; start++ {
		temp := twoSum(nums[start+1:], 0-nums[start])
		if len(temp) > 0 {
			for i := range temp {
				result = append(result, append(temp[i], nums[start]))
			}
		}
		for start < len(nums)-1 && nums[start] == nums[start+1] {
			start++
		}
	}

	return result
}

func twoSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 2 {
		return result
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			leftVal := nums[left]
			rightVal := nums[right]
			arr := []int{leftVal, nums[right]}
			result = append(result, arr)
			for left < right && nums[left] == leftVal {
				left++
			}
			for left < right && nums[right] == rightVal {
				right--
			}
		} else if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		}
	}
	return result
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := 0, len(nums)-1

	pivot := rand.Int() % len(nums)

	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	quickSort(nums[:left])
	quickSort(nums[left+1:])
	return nums
}
