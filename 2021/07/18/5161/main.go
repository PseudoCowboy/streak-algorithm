package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canBeTypedWords("leet code", "e"))
}

func canBeTypedWords(text string, brokenLetters string) int {
	ta := strings.Split(text, " ")
	n := len(ta)
	m := make(map[byte]int)
	for i := range brokenLetters {
		m[brokenLetters[i]]++
	}
	for i := range ta {
		for j := range ta[i] {
			if m[ta[i][j]] != 0 {
				n--
				break
			}
		}
	}

	return n
}
