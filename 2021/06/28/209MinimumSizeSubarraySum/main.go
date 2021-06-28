package main

func main() {
	minSubArrayLen(4, []int{1, 4, 4})
}

func minSubArrayLen(target int, nums []int) int {
	count := 0
	left, right := 0, 0
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		if sum >= target {
			if count == 0 {
				count = right - left + 1
			}
			for sum >= target {
				if right-left < count {
					count = right - left + 1
				}
				sum -= nums[left]
				left++
			}
		}
		right++
	}

	return count
}
