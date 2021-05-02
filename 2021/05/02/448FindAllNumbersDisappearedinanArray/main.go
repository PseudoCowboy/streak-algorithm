package main

func findDisappearedNumbers(nums []int) []int {
	arr := make([]int, len(nums))
	for i := range nums {
		arr[nums[i]-1]++
	}

	result := make([]int, 0)
	for i := range arr {
		if arr[i] == 0 {
			result = append(result, i+1)
		}
	}

	return result
}
