package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getLucky("zbax", 2))
}

func getLucky(s string, k int) int {
	temp := ""
	for i := range s {
		temp += strconv.Itoa(int(s[i] - 'a' + 1))
	}
	num := 0
	for i := 0; i < k; i++ {
		num = 0
		for len(temp) > 0 {
			a, _ := strconv.Atoi(temp[:1])
			num += a
			temp = temp[1:]
		}
		temp = strconv.Itoa(num)
	}
	return num
}
