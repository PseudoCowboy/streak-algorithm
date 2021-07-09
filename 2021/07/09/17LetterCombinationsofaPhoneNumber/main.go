package main

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	dfs(digits, &result, "")
	return result
}

func dfs(digits string, result *[]string, current string) {
	if len(digits) == 0 {
		*result = append(*result, current)
		return
	}

	for _, v := range m[digits[0]] {
		dfs(digits[1:], result, current+v)
	}

}
