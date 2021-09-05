package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

func countQuadruplets(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for t := k + 1; t < len(nums); t++ {
					if nums[i]+nums[j]+nums[k] == nums[t] {
						ans++
					}
				}
			}
		}
	}
	return ans
}
