package main

import "fmt"

func main() {
	fmt.Println(canSeePersonsCount([]int{5, 1, 2, 3, 10}))
}

func canSeePersonsCount(heights []int) []int {
	ans := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		count := 0
		pre := 0
		for j := i + 1; j < len(heights); j++ {
			if heights[j] > heights[i] {
				count++
				break
			}
			if heights[j] > pre {
				count++
				pre = heights[j]
			}
		}
		ans[i] = count
	}

	return ans
}
