package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findOriginalArray([]int{0, 0, 0, 0}))
}

func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				count++
			}
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return []int{}
	}
	m := make(map[int]int)
	for _, v := range changed {
		m[v]++
	}
	sort.Ints(changed)
	ans := make([]int, 0)
	for i := len(changed) - 1; i >= 0; i-- {
		if changed[i] == 0 {
			if m[0] > 1 {
				m[0] -= 2
				ans = append(ans, 0)
			}
			continue
		}
		if changed[i]%2 == 0 {
			v := changed[i] / 2
			if m[v] > 0 && m[changed[i]] > 0 {
				m[v]--
				m[changed[i]]--
				ans = append(ans, v)
			}
		}
	}
	if len(ans) == len(changed)/2 {
		return ans
	}
	return []int{}
}
