package main

func main() {
	getMaxLen([]int{1, 2, 3, 5, -6, 4, 0, 10})
}

func getMaxLen(nums []int) int {
	zeroCount := 0
	firstNeg := 0
	secondNeg := 0
	zeroIndex := 0
	length := 0
	nums = append(nums, 0)
	for i := range nums {
		if nums[i] == 0 {
			if zeroCount%2 == 0 {
				length = max(length, i-1-zeroIndex)
			} else {
				if zeroCount > 1 {
					length = max(length, max(max(i-1-firstNeg, firstNeg-zeroIndex), max(i-1-secondNeg, secondNeg-zeroIndex)))
				} else {
					length = max(length, max(i-1-firstNeg, firstNeg-zeroIndex))
				}
			}
			zeroIndex = i + 1
			firstNeg = i + 1
			zeroCount = 0
			continue
		}
		if nums[i] < 0 && zeroCount > 1 {
			secondNeg = i
			zeroCount++
		} else if nums[i] < 0 {
			firstNeg = i
			zeroCount++
		}

	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
