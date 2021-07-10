package main

func partition(s string) [][]string {
	result := make([][]string, 0)
	dfs(s, []string{}, &result)
	return result
}

func dfs(s string, current []string, result *[][]string) {
	if len(s) == 0 {
		temp := make([]string, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			current = append(current, s[:i])
			dfs(s[i:], current, result)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
