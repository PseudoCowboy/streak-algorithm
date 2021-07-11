package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("0000"))
}

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n > 12 || n < 4 {
		return []string{}
	}
	result := make([]string, 0)
	dfs(&result, []string{}, s, 1)
	return result
}

func dfs(result *[]string, current []string, s string, index int) {
	if len(s) == 0 {
		temp := strings.Join(current, ".")
		*result = append(*result, temp)
		return
	}
	if index == 5 {
		return
	}

	for i := 0; i < 3; i++ {
		if i+1 > len(s) {
			continue
		}
		v, _ := strconv.Atoi(s[:i+1])
		if len(s[i+1:]) > 3*(4-index) || len(s[i+1:]) < 4-index || (len(s[:i+1]) > 1 && s[:i+1][0] == '0') || v > 255 {
			continue
		}
		current = append(current, s[:i+1])
		dfs(result, current, s[i+1:], index+1)
		current = current[:len(current)-1]
	}
}
