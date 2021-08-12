package main

import "fmt"

func main() {
	fmt.Println(lastSubstring("leetcode"))
}

func lastSubstring(s string) string {
	arr := make([]int, 26)
	for i := range arr {
		arr[i] = -1
	}
	for i := len(s) - 1; i >= 0; i-- {
		arr[s[i]-'a'] = i
	}
	target := 0
	for i := 25; i >= 0; i-- {
		if arr[i] != -1 {
			target = arr[i]
			break
		}
	}
	return s[target:]
}
