package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				fmt.Println(i, nums1[i], nums2[n-1])
				n--
			}
			continue
		}
		if m > 0 {
			nums1[i] = nums1[m-1]
			m--
		}
		if n > 0 {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}
