package main

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		s1 := find(s, i, i)
		s2 := find(s, i, i+1)
		if len(s1) > len(ans) {
			ans = s1
		}
		if len(s2) > len(ans) {
			ans = s2
		}
	}
	return ans
}

func find(s string, start, end int) string {
	sub := ""
	for start >= 0 && end < len(s) && s[start] == s[end] {
		sub = s[start : end+1]
		start--
		end++
	}
	return sub
}
