package main

func main() {
	missingNumber([]int{3, 0, 1})
}

func missingNumber(nums []int) int {
	result := 0
	for i := range nums {
		result = result ^ i ^ nums[i]
	}
	return result ^ len(nums)
}
