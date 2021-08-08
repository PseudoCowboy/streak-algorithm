package main

import "fmt"

func main() {
	partition("aab")
}

func partition(s string) [][]string {
	ans := make([][]string, 0)
	dfs(&ans, []string{}, s)
	return ans
}

func dfs(ans *[][]string, current []string, s string) {
	if len(s) == 0 {
		fmt.Println(current)
		temp := make([]string, len(current))
		copy(temp, current)
		*ans = append(*ans, temp)
		return
	}

	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			current = append(current, s[:i])
			dfs(ans, current, s[i:])
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
