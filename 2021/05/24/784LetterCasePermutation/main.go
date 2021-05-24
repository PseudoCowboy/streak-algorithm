package main

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	dfs(&result, s, 0, "")
	return result
}

func dfs(result *[]string, s string, index int, current string) {
	if len(s) == index {
		*result = append(*result, current)
		return
	}

	if s[index] >= 'A' && s[index] <= 'Z' {
		dfs(result, s, index+1, current+string(s[index]+byte(32)))
	}
	if s[index] >= 'a' && s[index] <= 'z' {
		dfs(result, s, index+1, current+string(s[index]-byte(32)))
	}
	dfs(result, s, index+1, current+string(s[index]))
}
