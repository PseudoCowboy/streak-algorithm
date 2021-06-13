package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumRemovals("abcab", "abc", []int{0, 1, 2, 3, 4}))
}

func maximumRemovals(s string, p string, removable []int) int {
	max := 0
	index := 0
	for index < len(s) {
		i := 0
		j := 0
		for i < len(s) && j < len(p) && index+1 <= len(removable) {
			temp := removable[:index+1]
			sort.Ints(temp)
			check := isInRemovable(temp, i)
			if check != -1 {
				i = temp[check] + 1
				continue
			}
			if s[i] == p[j] {
				j++
			}
			i++
		}
		if j != len(p) {
			return max
		}
		index++
		max = index
	}
	return max
}

func isInRemovable(arr []int, target int) int {
	for i := range arr {
		if target == arr[i] {
			return i
		}
	}
	return -1
}
