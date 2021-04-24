package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 0
	for i++; i < 10; i++ {
		fmt.Println(i)
	}

	// fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	for i, n := 0, len(nums); i < n; {
		left := i
		for ; i < n && nums[i-1]+1 == nums[i]; i++ {
			fmt.Println(i)
		}
		s := strconv.Itoa(nums[left])
		if left != i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		result = append(result, s)
	}
	return result
}
