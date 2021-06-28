package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	outer := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[outer] = nums[outer], nums[i]
			outer--
		}
	}

	return outer + 1
}
