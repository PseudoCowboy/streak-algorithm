package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	arr := make([]int, 0)
	for i := range nums1 {
		_, ok := m[nums1[i]]
		if ok {
			m[nums1[i]]++
		} else {
			m[nums1[i]] = 1
		}
	}

	for i := range nums2 {
		val, ok := m[nums2[i]]
		if ok {
			if val >= 1 {
				arr = append(arr, nums2[i])
				m[nums2[i]]--
			}
		}
	}

	return arr
}
