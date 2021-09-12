package main

import "strings"

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "") + ")"
	ans, _ := calc(s, 0)
	return ans
}

func calc(s string, startIndex int) (int, int) {
	ans := 0
	current := 0
	sign := 1
	for i := startIndex; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			current = current*10 + int(s[i]-'0')
			continue
		}

		if s[i] == '(' {
			inner, nextStart := calc(s, i+1)
			ans += sign * inner
			i = nextStart
		} else if s[i] == ')' {
			return ans, i
		} else if s[i] == '+' {
			sign = 1
		} else if s[i] == '-' {
			sign = -1
		}
	}

	return ans, len(s)
}
