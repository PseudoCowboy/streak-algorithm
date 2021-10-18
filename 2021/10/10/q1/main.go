package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	m3 := make(map[int]int)
	for i := range nums1 {
		m1[nums1[i]]++
	}
	for i := range nums2 {
		m2[nums2[i]]++
	}
	for i := range nums3 {
		m3[nums3[i]]++
	}
	m := make(map[int]int)
	ans := make([]int, 0)
	for i := range nums1 {
		if m2[nums1[i]] > 0 {
			m[nums1[i]]++
		}
		if m3[nums1[i]] > 0 {
			m[nums1[i]]++
		}
	}
	for i := range nums2 {
		if m3[nums2[i]] > 0 {
			m[nums2[i]]++
		}
	}
	for k := range m {
		ans = append(ans, k)
	}
	return ans
}
