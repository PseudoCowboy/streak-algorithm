package main

import "fmt"

func main() {
	binarySearch2([]int{1, 2, 2, 2, 2, 3, 5}, 0)
}

func binarySearch2(arr []int, target int) {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] == target {
			left = mid + 1
		}
		if arr[mid] < target {
			left = mid + 1
		}
		if arr[mid] > target {
			right = mid
		}
	}

	fmt.Println(left, right)
}

func binarySearch1(arr []int, target int) {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] == target {
			left = mid
		}
		if arr[mid] < target {
			left = mid + 1
		}
		if arr[mid] > target {
			right = mid
		}
	}

	fmt.Println(left, right)
}

func binarySearch(arr []int, target int) {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == target {
			right = mid - 1
		}
		if arr[mid] < target {
			left = mid + 1
		}
		if arr[mid] > target {
			right = mid - 1
		}
	}

	fmt.Println(left, right)
}
