package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(countValidWords("a-!d"))
}

func countValidWords(sentence string) int {
	ss := strings.Fields(sentence)
	ans := 0
	for i := range ss {
		if len(ss[i]) == 1 && (ss[i] == "," || ss[i] == "." || ss[i] == "!") {
			ans++
			continue
		}
		if (ss[i][0] >= 'a' && ss[i][0] <= 'z') && ((ss[i][len(ss[i])-1] >= 'a' && ss[i][len(ss[i])-1] <= 'z') || ss[i][len(ss[i])-1] == '.' || ss[i][len(ss[i])-1] == ',' || ss[i][len(ss[i])-1] == '!') {
			a := 0
			b := 0
			c := 0
			for j := 1; j < len(ss[i])-1; j++ {
				if ss[i][j] == '-' {
					a++
				}
				if ss[i][j] >= '0' && ss[i][j] <= '9' {
					b++
				}
				if ss[i] == "," || ss[i] == "." || ss[i] == "!" {
					c++
				}
			}
			if a < 2 && b < 1 && c < 1 {
				ans++
			}
		}
	}
	return ans
}
