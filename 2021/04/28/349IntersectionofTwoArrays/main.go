package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := range nums1 {
		m[nums1[i]] = 0
	}

	for i := range nums2 {
		_, ok := m[nums2[i]]
		if ok {
			m[nums2[i]]++
		}
	}

	result := make([]int, 0)
	for k, v := range m {
		if v > 0 {
			result = append(result, k)
		}
	}

	return result
}

func intersection1(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}
