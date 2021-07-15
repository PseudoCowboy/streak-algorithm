package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{0, 1, 2, 1, 2}))
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	mountain := 0
	downflag := false

	if arr[0] > arr[1] {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if arr[i] > arr[i-1] {
			downflag = false
		}
		if arr[i] < arr[i-1] {
			if !downflag {
				mountain++
			}
			downflag = true
		}
	}

	return downflag && mountain == 1
}
