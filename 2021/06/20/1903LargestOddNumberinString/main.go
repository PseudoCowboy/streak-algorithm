package main

func largestOddNumber(num string) string {
	if len(num) == 0 {
		return ""
	}
	n := len(num) - 1
	for n >= 0 {
		if num[n] == '1' || num[n] == '3' || num[n] == '5' || num[n] == '7' || num[n] == '9' {
			break
		}
		n--
	}
	return num[:n+1]
}
