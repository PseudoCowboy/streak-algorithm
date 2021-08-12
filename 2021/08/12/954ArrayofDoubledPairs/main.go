package main

import "sort"

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	count := 0
	m := make(map[int]int)
	for i := range arr {
		if val, ok := m[arr[i]*2]; ok {
			if val > 0 {
				m[arr[i]*2]--
				count++
				continue
			}
		}
		if arr[i]%2 == 0 {
			if val, ok := m[arr[i]/2]; ok {
				if val > 0 {
					m[arr[i]/2]--
					count++
					continue
				}
			}
		}
		m[arr[i]]++
	}
	return count == len(arr)/2
}

func canReorderDoubled1(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) > abs(arr[j])
	})
	count := 0
	m := make(map[int]int)
	for i := range arr {
		if val, ok := m[arr[i]*2]; ok {
			if val > 0 {
				m[arr[i]*2]--
				count++
				continue
			}
		}
		m[arr[i]]++
	}
	return count == len(arr)/2
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
