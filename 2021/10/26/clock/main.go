package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(2, 3, 6, 0))
}

func Solution(A int, B int, C int, D int) int {
	// write your code in Go 1.4
	digits := []int{A, B, C, D}
	used := make([]bool, 4)
	ans := make(map[string]int)
	dfs(used, digits, "", ans)
	fmt.Println(ans)
	return len(ans)
}

func dfs(used []bool, digits []int, current string, ans map[string]int) {
	if len(current) == 4 {
		ans[current]++
	}

	for i := 0; i < len(digits); i++ {
		if !used[i] {
			switch len(current) {
			case 0:
				if !(digits[i] <= 2) {
					continue
				}
			case 1:
				if !(current[0] == '0' || current[0] == '1' || (current[0] == '2' && digits[i] <= 4)) {
					continue
				}
			case 2:
				if !(digits[i] <= 5) {
					continue
				}
			default:
			}
			currentDigit := strconv.Itoa(digits[i])
			current += currentDigit
			used[i] = true
			dfs(used, digits, current, ans)
			used[i] = false
			current = current[:len(current)-1]
		}
	}
}
