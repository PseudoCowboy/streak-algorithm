package main

import "fmt"

func main() {
	fmt.Println(isValid("(("))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := map[byte]byte{40: 41, 123: 125, 91: 93}
	stack := []byte{}
	for i := range s {
		_, ok := m[s[i]]
		if ok {
			stack = append(stack, s[i])
		} else {
			stackLenth := len(stack)
			if stackLenth == 0 || s[i]-stack[stackLenth-1] > 2 || s[i]-stack[stackLenth-1] < 1 {
				return false
			} else {
				stack = stack[:stackLenth-1]
			}
		}
	}
	return len(stack) == 0
}
