package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
		} else {
			if temp > max {
				max = temp
			}
			temp = 0
		}
	}
	if temp > max {
		return temp
	}
	return max
}
