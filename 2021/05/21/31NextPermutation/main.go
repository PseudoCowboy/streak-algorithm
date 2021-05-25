package main

import "fmt"

func main() {
	arr := []int{4, 7, 5, 3, 2, 1}
	i := 0
	for i = len(arr) - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			break
		}
	}
	fmt.Println(i)
}

func nextPermutation(nums []int) {
	i := 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums)
		return
	}
	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j]
			break
		}
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
