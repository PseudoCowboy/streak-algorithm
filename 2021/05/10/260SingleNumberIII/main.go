package main

import "fmt"

func main() {
	a := 8
	b := 11
	result := 0 ^ a ^ b

	result &= -result
	fmt.Println(result)
}

func singleNumber(nums []int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	result := make([]int, 0)
	for k, v := range m {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}

func singleNumber1(nums []int) []int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	result &= -result
	arr := []int{0, 0}
	for _, v := range nums {
		if result&v == 0 {
			arr[0] ^= v
		} else {
			arr[1] ^= v
		}
	}
	return arr
}
