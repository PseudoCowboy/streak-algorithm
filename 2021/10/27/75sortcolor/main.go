package main

import "fmt"

func main() {
	sortColors([]int{0, 0, 0, 2, 2, 0, 2, 2, 0, 0, 0, 2, 2, 0, 2, 2})
}

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == 0 {
			left++
			continue
		}
		if nums[left] == 2 {
			for right > left {
				if nums[right] == 0 {
					break
				}
				right--
			}
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
	}

	fmt.Println(nums)
}
