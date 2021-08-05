package main

func stoneGame(piles []int) bool {
	a := 0
	l := 0
	mid := len(piles) / 2
	left := piles[:mid]
	reverse(left)
	// right := piles[mid:]
	// dp := make([]int, mid)

	return a > l
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
