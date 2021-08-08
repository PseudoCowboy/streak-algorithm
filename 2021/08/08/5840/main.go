package main

import "fmt"

func main() {
	fmt.Println(minSwaps("]]][[["))
}

func minSwaps(s string) int {
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	count := 0
	for i := range stack {
		if stack[i] == ']' {
			count++
		}
	}
	if count%2 == 1 {
		count++
	}
	return count / 2
}
