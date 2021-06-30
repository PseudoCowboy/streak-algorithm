package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abc", "cba"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := [26]int{}
	for i := range s {
		arr[s[i]-'a']++
	}
	for i := range t {
		arr[t[i]-'a']--
	}

	for i := range arr {
		if arr[i] != 0 {
			return false
		}
	}

	return true
}
