package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	sm := make(map[byte]byte)
	tm := make(map[byte]byte)
	for i := range s {
		sv, sok := sm[s[i]]
		tv, tok := tm[t[i]]
		if !sok && !tok {
			sm[s[i]] = t[i]
			tm[t[i]] = s[i]
			continue
		}
		if !sok || !tok {
			return false
		}
		if sv != t[i] || tv != s[i] {
			return false
		}
	}

	return true
}
