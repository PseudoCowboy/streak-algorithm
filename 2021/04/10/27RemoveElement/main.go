package main

func main() {
	removeElement([]int{3, 3}, 3)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	outer := len(nums) - 1
	for i := outer; i >= 0; i-- {
		if nums[i] == val {
			nums[i] = nums[outer]
			outer--
		}
	}
	return outer + 1
}
