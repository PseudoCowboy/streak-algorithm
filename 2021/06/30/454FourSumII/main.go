package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for i := range nums1 {
		for j := range nums2 {
			m[nums1[i]+nums2[j]]++
		}
	}
	count := 0
	for i := range nums3 {
		for j := range nums4 {
			count += m[-nums3[i]-nums4[j]]
		}
	}

	return count
}
