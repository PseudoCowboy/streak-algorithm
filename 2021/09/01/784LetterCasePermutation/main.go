package main

func letterCasePermutation(s string) []string {
	ans := make([]string, 0)
	dfs(s, []byte{}, &ans)
	return ans
}

func dfs(s string, current []byte, ans *[]string) {
	if len(current) == len(s) {
		*ans = append(*ans, string(current))
		return
	}

	n := len(current)
	dfs(s, append(current, s[n]), ans)
	if s[n] >= 'a' && s[n] <= 'z' {
		dfs(s, append(current, s[n]-32), ans)
	}
	if s[n] >= 'A' && s[n] <= 'Z' {
		dfs(s, append(current, s[n]+32), ans)
	}
}
