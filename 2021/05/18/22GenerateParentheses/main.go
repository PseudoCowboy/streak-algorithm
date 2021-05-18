package main

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	if n == 0 {
		return result
	}
	dfs(&result, n, n, "")

	return result
}

func dfs(result *[]string, left, right int, current string) {
	if left == 0 && right == 0 {
		*result = append(*result, current)
		return
	}

	if left > 0 {
		dfs(result, left-1, right, current+"(")
	}
	if right > 0 && left < right {
		dfs(result, left, right-1, current+")")
	}

}
