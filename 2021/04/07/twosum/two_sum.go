package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]] = i

		num, ok := m[target-nums[i]]
		if ok {
			return []int{num, i}
		}
	}

	return []int{}
}
