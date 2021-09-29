package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 3, 2, 1}
	sort.Ints(arr[:3])
	fmt.Println(arr)
}

func sortArrayByParityII(nums []int) []int {
	odd := 1
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			continue
		}
		for nums[odd]%2 == 1 {
			odd += 2
		}
		nums[i], nums[odd] = nums[odd], nums[i]
	}

	return nums
}
