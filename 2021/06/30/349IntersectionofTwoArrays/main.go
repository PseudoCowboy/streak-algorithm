package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m := make(map[int]int)

	for i := range nums1 {
		m[nums1[i]] = 1
	}
	ans := make([]int, 0)
	for i := range nums2 {
		v := m[nums2[i]]
		if v > 0 {
			ans = append(ans, nums2[i])
			m[nums2[i]]--
		}
	}

	return ans
}
