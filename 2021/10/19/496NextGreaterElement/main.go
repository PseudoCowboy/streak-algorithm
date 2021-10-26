package main

import (
	"fmt"
	"strconv"
)

func main() {
	nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7})
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums2) == 1 {
		return []int{-1}
	}
	res := make([]int, len(nums2))
	m := make(map[int]int)
	for i := range nums2 {
		m[nums2[i]] = i
	}
	stack := make([]int, 0)
	stack = append(stack, len(nums2)-1)
	res[len(nums2)-1] = -1
	for i := len(nums2) - 2; i >= 0; i-- {
		fmt.Println(stack, nums2[i], stack[len(stack)-1], nums2[stack[len(stack)-1]])
		if nums2[i] < nums2[stack[len(stack)-1]] {
			res[i] = nums2[stack[len(stack)-1]]
			stack = append(stack, i)
		}
		if nums2[i] > nums2[stack[len(stack)-1]] {
			for len(stack) != 0 && nums2[i] > nums2[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				res[i] = -1
			} else {
				res[i] = nums2[stack[len(stack)-1]]
			}
			stack = append(stack, i)
		}
	}
	for i := range nums1 {
		nums1[i] = res[m[nums1[i]]]
	}
	return nums1
}

func Solution(A int, B int, C int, D int) int {
	// write your code in Go 1.4
	ans := 0
	digits := []int{A, B, C, D}
	used := make([]bool, 4)
	dfs(used, digits, 0, "", &ans)
	return ans
}

func dfs(used []bool, digits []int, index int, current string, ans *int) {
	if len(current) == 4 {
		*ans++
	}

	for i := index; i < len(digits); i++ {
		if !used[i] {
			switch len(current) {
			case 0:
				if digits[i] <= 2 {
					currentDigit := strconv.Itoa(digits[i])
					current += currentDigit
					used[i] = true
					dfs(used, digits, i+1, current, ans)
					used[i] = false
					current = current[:len(current)-1]
				}
			case 1:
				if current[0] == '0' || current[0] == '1' || (current[0] == '2' && digits[i] <= 4) {
					currentDigit := strconv.Itoa(digits[i])
					current += currentDigit
					used[i] = true
					dfs(used, digits, i+1, current, ans)
					used[i] = false
					current = current[:len(current)-1]
				}
			case 2:
				if digits[i] <= 5 {
					currentDigit := strconv.Itoa(digits[i])
					current += currentDigit
					used[i] = true
					dfs(used, digits, i+1, current, ans)
					used[i] = false
					current = current[:len(current)-1]
				}
			default:
				currentDigit := strconv.Itoa(digits[i])
				current += currentDigit
				used[i] = true
				dfs(used, digits, i+1, current, ans)
				used[i] = false
				current = current[:len(current)-1]
			}
		}
	}

}
