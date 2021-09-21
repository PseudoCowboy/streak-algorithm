package main

import "strings"

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, v := range operations {
		if strings.Contains(v, "++") {
			ans++
		} else {
			ans--
		}
	}

	return ans

}
