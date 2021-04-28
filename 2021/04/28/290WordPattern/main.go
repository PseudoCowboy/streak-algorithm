package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")
	if len(sArr) != len(pattern) {
		return false
	}
	patternMap := make(map[byte]string)
	sMap := make(map[string]byte)
	for i := 0; i < len(sArr); i++ {
		fmt.Println(pattern[i], sArr[i])
		pVal, pOk := patternMap[pattern[i]]
		sVal, sOk := sMap[sArr[i]]

		if (!pOk && sOk) || (!sOk && pOk) {
			return false
		}
		if pOk && sOk {
			fmt.Println(pVal, sVal, sArr[i], pattern[i])
			if pVal != sArr[i] || sVal != pattern[i] {
				return false
			}
		} else {
			patternMap[pattern[i]] = sArr[i]
			sMap[sArr[i]] = pattern[i]
		}
	}

	return true

}
