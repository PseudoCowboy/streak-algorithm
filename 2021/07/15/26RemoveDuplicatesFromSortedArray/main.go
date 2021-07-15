package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != pre {
			pre = nums[i]
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	return count
}
