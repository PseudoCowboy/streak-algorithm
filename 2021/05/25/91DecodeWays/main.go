package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	start, result := 0, 0
	dfs(s, start, &result)
	return result
}

func dfs(s string, start int, result *int) {
	if start == len(s) {
		*result++
		return
	}

	if s[start] == '1' {
		if start+2 <= len(s) {
			dfs(s, start+2, result)
		}
	}
	if s[start] == '2' {
		if start+2 <= len(s) && s[start+1] <= '6' {
			dfs(s, start+2, result)
		}
	}
	if s[start+1] != '0' {
		dfs(s, start+1, result)
	}
}
