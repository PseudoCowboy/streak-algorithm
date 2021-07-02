package main

import "fmt"

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {
	for i := 2; i <= len(s); i++ {
		if len(s)%i != 0 {
			continue
		}
		subSize := len(s) / i
		flag := true
		for j := 0; j < subSize; j++ {
			for k := 1; k < i; k++ {
				if s[j] != s[k*subSize+j] {
					flag = false
				}
			}
		}
		if flag == true {
			return true
		}
	}

	return false
}
