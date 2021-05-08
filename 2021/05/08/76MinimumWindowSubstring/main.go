package main

import "math"

func minWindow(s string, t string) string {
	window := make(map[byte]int)
	need := make(map[byte]int)
	valid := 0
	left, right, start := 0, 0, 0
	length := math.MaxInt16

	for i := range t {
		need[t[i]]++
	}

	for right < len(s) {
		if val, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == val {
				valid++
			}
		}
		right++

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			if val, ok := need[s[left]]; ok {
				if val == window[s[left]] {
					valid--
				}
				window[s[left]]--
			}

			left++
		}
	}

	if length == math.MaxInt16 {
		return ""
	}
	return s[start : start+length]
}
