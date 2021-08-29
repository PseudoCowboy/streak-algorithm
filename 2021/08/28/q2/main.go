package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("0" > "0")

	arr := []string{"1", "2", "12", "21"}
	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) == len(arr[j]) {
			return arr[i] > arr[j]
		}
		return len(arr[i]) > len(arr[j])
	})
	fmt.Println(arr)
	fmt.Println(arr[2])
}

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) == len(nums[j]) {
			return nums[i] > nums[j]
		}
		return len(nums[i]) > len(nums[j])
	})
	return nums[k-1]
}
