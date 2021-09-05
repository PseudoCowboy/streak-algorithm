package main

import "sort"

func orderlyQueue(s string, k int) string {
	if k == 1 {
		n := len(s)
		temp := s
		s += s
		for i := 1; i < n; i++ {
			if s[i:i+n] < temp {
				temp = s[i : i+n]
			}
		}
		return temp
	}
	sb := []byte(s)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	return string(sb)
}
