package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	aStr := []byte(s)
	bStr := []byte(t)

	aMap := make(map[byte]byte)
	bMap := make(map[byte]byte)

	for i, val := range aStr {
		if _, ok := aMap[val]; ok {
			if bVal, ok := bMap[bStr[i]]; ok {
				if bVal != val {
					return false
				}
			} else {
				return false
			}
		} else {
			if _, ok := bMap[bStr[i]]; ok {
				return false
			} else {
				aMap[val] = bStr[i]
				bMap[bStr[i]] = val
			}
		}
	}
	return true
}
