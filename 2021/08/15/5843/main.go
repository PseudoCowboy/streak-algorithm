package main

import "strings"

func numOfStrings(patterns []string, word string) int {
	count := 0
	for i := range patterns {
		if strings.Contains(word, patterns[i]) {
			count++
		}
	}
	return count
}
