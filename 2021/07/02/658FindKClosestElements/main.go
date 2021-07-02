package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
}

func findClosestElements1(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1-k
	for left <= right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return arr[left : left+k]
}

func findClosestElements(arr []int, k int, x int) []int {
	index := findIndex(arr, x)
	ans := make([]int, 0)

	left, right := index, index+1
	for k != 0 {
		if left >= 0 && right <= len(arr)-1 {
			if x-arr[left] <= arr[right]-x {
				ans = append(ans, arr[left])
				left--
			} else {
				ans = append(ans, arr[right])
				right++
			}
		} else if left >= 0 {
			ans = append(ans, arr[left])
			left--

		} else {
			ans = append(ans, arr[right])
			right++
		}
		k--
	}

	sort.Ints(ans)
	return ans
}

func findIndex(arr []int, x int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			right = mid - 1
		} else if arr[mid] < x {
			left = mid + 1
		}
	}
	if right < 0 {
		right = 0
	}
	return right
}
