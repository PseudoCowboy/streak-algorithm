package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{0, 1, 2}))
}

func sortArrayByParity(nums []int) []int {
	odd := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			continue
		}
		for nums[odd]%2 == 1 {
			odd--
			if odd < i {
				return nums
			}
		}
		nums[i], nums[odd] = nums[odd], nums[i]
	}
	return nums
}
