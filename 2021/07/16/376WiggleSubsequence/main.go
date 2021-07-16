package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{0, 0}))
}

func wiggleMaxLength(nums []int) int {
	count := 1
	preD := 0
	curD := 0
	for i := 1; i < len(nums); i++ {
		curD = nums[i] - nums[i-1]
		if (curD > 0 && preD <= 0) || (curD < 0 && preD >= 0) {
			preD = curD
			count++
		}
	}

	return count
}
