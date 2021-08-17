package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	max := 0
	m := make(map[byte]int)
	for right < len(s) {
		m[s[right]]++
		for m[s[right]] > 1 {
			m[s[left]]--
			left++
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		right++
	}
	return max
}
