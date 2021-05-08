package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := (len(nums1) + len(nums2)) / 2
	first, second := 0, 0
	for first+second < mid-1 {
		if nums1[first] < nums2[second] {
			if first == len(nums1) {
				second++
			} else {
				first++
			}
		} else {
			if second == len(nums2) {
				second++
			} else {
				first++
			}
		}
	}

	return 0.0
}
