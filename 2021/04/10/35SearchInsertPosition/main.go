package main

func main() {

}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || nums[0] >= target {
		return 0
	}
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
