package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumNumber("214010", []int{6, 7, 9, 7, 4, 0, 3, 4, 4, 7}))
}

func maximumNumber1(num string, change []int) string {
	ans := ""
	flag := true
	for i := range num {
		index := int(num[i] - '0')
		if index < change[index] {
			ans += strconv.Itoa(change[index])
			flag = false
		} else {
			if !flag {
				ans += num[i:]
				return ans
			}
			ans += num[i : i+1]
		}
	}

	return ans
}

// "214010"
// [6,7,9,7,4,0,3,4,4,7]
func maximumNumber(num string, change []int) string {
	ans := make([]byte, len(num))
	flag := true
	for i := range num {
		index := num[i] - '0'
		if index <= byte(change[index]) {
			ans[i] = byte(change[index]) + '0'
			if index != byte(change[index]) {
				flag = false
			}
		} else {
			if !flag {
				return string(ans[:i]) + num[i:]
			}
			ans[i] += num[i]
		}
	}

	return string(ans)
}
