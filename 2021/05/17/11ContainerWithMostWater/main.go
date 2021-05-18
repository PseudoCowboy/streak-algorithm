package main

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(h)
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		tempHeight := 0
		leftlow := false
		if height[left] > height[right] {
			leftlow = false
			tempHeight = height[right]
		} else {
			leftlow = true
			tempHeight = height[left]
		}
		area := tempHeight * (right - left)
		if area > max {
			max = area
		}
		if leftlow {
			left++
		} else {
			right--
		}
	}
	return max
}
