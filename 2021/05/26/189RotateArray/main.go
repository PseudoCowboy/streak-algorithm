package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
