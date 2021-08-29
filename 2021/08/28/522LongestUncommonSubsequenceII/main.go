package main

func main() {
	findLUSlength([]string{"aaa", "aaa", "aa"})
}

func findLUSlength(strs []string) int {
	ans := -1
	for i := 0; i < len(strs); i++ {
		isUnCommon := true
		for j := 0; j < len(strs); j++ {
			if i != j && isValid(strs[i], strs[j]) {
				isUnCommon = false
				break
			}
		}
		if isUnCommon {
			ans = max(ans, len(strs[i]))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValid(a, b string) bool {
	if len(a) > len(b) {
		return false
	}

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}
