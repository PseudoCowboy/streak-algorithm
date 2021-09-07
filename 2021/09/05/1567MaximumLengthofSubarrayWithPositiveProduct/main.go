package main

import "fmt"

func main() {
	getMaxLen([]int{1, -2, -3, 4})
}

func getMaxLen(nums []int) int {
	if len(nums) == 1 {
		if nums[0] > 1 {
			return 1
		}
		return 0
	}
	arr := make([]int, len(nums))
	if nums[0] > 0 {
		arr[0] = 1
	}
	if nums[0] == 0 {
		arr[0] = 0
	}
	if nums[0] < 0 {
		arr[0] = -1
	}
	for i := 1; i < len(nums); i++ {
		if arr[i-1] == 0 {
			arr[i] = check(nums[i])
			continue
		}
		arr[i] = arr[i-1] * check(nums[i])
	}
	fmt.Println(arr)
	sum := 0
	temp := 0
	for i := range arr {
		if arr[i] > 0 {
			temp++
			if temp > sum {
				sum = temp
			}
		} else {
			temp = 0
		}
	}
	return sum
}

func check(a int) int {
	if a < 0 {
		return -1
	}
	if a == 0 {
		return 0
	}
	return 1
}
