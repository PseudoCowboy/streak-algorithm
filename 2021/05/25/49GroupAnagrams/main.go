package main

import (
	"math/rand"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		key := make([]byte, len(v))
		for i := range v {
			key[i] = v[i]
		}
		quickSort(key)
		if val, ok := m[string(key)]; ok {
			m[string(key)] = append(val, v)
		} else {
			m[string(key)] = []string{v}
		}
	}
	result := make([][]string, 0)
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func quickSort(nums []byte) {
	if len(nums) < 2 {
		return
	}

	left, right := 0, len(nums)-1
	pivot := rand.Int() % len(nums)
	nums[right], nums[pivot] = nums[pivot], nums[right]
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[right], nums[left] = nums[left], nums[right]
	quickSort(nums[:left])
	quickSort(nums[left+1:])

}
