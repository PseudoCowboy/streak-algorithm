package main

import "fmt"

func main() {
	fmt.Println(isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"}))
	fmt.Println(9 / 2)
}

func isPrefixString(s string, words []string) bool {
	temp := ""
	n := len(s)
	for i := range words {
		temp += words[i]
		if len(temp) == n {
			if temp == s {
				return true
			} else {
				return false
			}
		}
		if len(temp) > n {
			return false
		}
	}

	return false
}
