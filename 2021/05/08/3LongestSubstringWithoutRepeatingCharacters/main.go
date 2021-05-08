package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	length := 0
	left, right := 0, 0
	for right < len(s) {
		t := s[right]
		m[t]++
		right++
		for m[t] > 1 {
			m[s[left]]--
			left++
		}
		if right-left > length {
			length = right - left
		}
	}

	return length
}
