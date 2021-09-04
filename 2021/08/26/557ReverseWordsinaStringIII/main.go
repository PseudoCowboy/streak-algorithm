package main

import "strings"

func reverseWords(s string) string {
	start := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			start = s[:i]
			break
		}
	}
	end := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = s[i+1:]
			break
		}
	}
	sa := strings.Split(s, " ")
	for i := range sa {
		sa[i] = reverse(sa[i])
	}
	return start + strings.Join(sa, " ") + end

}

func reverse(s string) string {
	arr := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return string(arr)
}
