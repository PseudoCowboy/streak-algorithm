package main

import "strings"

func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse(ss)
	return strings.Join(ss, " ")
}

func reverse(sArr []string) {
	left, right := 0, len(sArr)-1
	for left < right {
		sArr[left], sArr[right] = sArr[right], sArr[left]
		left++
		right--
	}
}
