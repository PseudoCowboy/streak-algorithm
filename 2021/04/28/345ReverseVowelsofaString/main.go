package main

import (
	"fmt"
)

func main() {
	s := "aeiou"
	arr := []byte{97, 101, 105, 111, 117}

	for i := range s {
		fmt.Println(s[i] == arr[0])
	}
}

func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if isVowels(b[i]) && isVowels(b[j]) {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		} else if isVowels(b[i]) && !isVowels(b[j]) {
			j--
		} else if !isVowels(b[i]) && isVowels(b[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	return string(b)
}

func isVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
