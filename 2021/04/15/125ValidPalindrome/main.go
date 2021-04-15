package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start, end := 0, len(s)-1

	for start < end {
		for start < end && !isChar(s[start]) {
			start++
		}
		for start < end && !isChar(s[end]) {
			end--
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func isChar(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}
