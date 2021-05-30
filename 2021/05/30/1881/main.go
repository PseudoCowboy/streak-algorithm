package main

import (
	"strconv"
)

func maxValue(n string, x int) string {
	index := 0
	if n[0] == '-' {
		for i := 1; i < len(n); i++ {
			if int(n[i]-'0') > x {
				index = i
				break
			}
			if i == len(n)-1 {
				index = len(n)
			}
		}

	} else {
		for i := 0; i < len(n); i++ {
			if int(n[i]-'0') < x {
				index = i
				break
			}
			if i == len(n)-1 {
				index = len(n)
			}
		}
	}

	return n[:index] + strconv.Itoa(x) + n[index:]
}
