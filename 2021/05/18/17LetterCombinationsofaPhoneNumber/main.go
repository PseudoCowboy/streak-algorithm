package main

var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	dfs(digits, &res, "")
	return res
}

func dfs(digits string, res *[]string, current string) {
	if digits == "" {
		*res = append(*res, current)
		return
	}
	k := digits[0:1]
	for i := 0; i < len(dict[k]); i++ {
		current += dict[k][i]
		dfs(digits[1:], res, current)
		current = current[:len(current)-1]
	}
}
