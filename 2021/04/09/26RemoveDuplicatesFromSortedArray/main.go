package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3, 3}))
}

func removeDuplicates(nums []int) int {
	index := 0
	if len(nums) == 0 {
		return index
	}
	temp := nums[0]
	for i := range nums {
		if nums[i] != temp {
			index++
			nums[index] = nums[i]
			temp = nums[i]
		}
	}
	return index + 1
}
