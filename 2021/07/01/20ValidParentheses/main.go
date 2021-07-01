package main

import "fmt"

func main() {
	fmt.Println(isValid("(("))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := map[byte]byte{
		125: 123,
		93:  91,
		41:  40,
	}

	stack := make([]byte, 0)

	for i := range s {
		if v, ok := m[s[i]]; ok {
			if len(stack) < 1 {
				return false
			}
			if v != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
