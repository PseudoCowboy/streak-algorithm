package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow1("aa", "aa"))
}

func minWindow1(s string, t string) string {
	need := make(map[byte]int)
	win := make(map[byte]int)
	valid, left, right := 0, 0, 0
	result := ""
	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		if val, ok := need[s[right]]; ok {
			win[s[right]]++
			if win[s[right]] == val {
				valid++
			}
		}
		right++
		for valid == len(need) {
			if left < right {
				if val, ok := need[s[left]]; ok {
					if val == win[s[left]] {
						fmt.Println(val)
						if result == "" || len(result) > right-left {
							result = s[left:right]
						}
						valid--
					}
					win[s[left]]--
				}
				left++
			}
		}
	}

	return result
}

func minWindow(s string, t string) string {
	window := make(map[byte]int)
	need := make(map[byte]int)
	valid := 0
	left, right, start := 0, 0, 0
	length := math.MaxInt16

	for i := range t {
		need[t[i]]++
	}

	for right < len(s) {
		if val, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == val {
				valid++
			}
		}
		right++

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			if val, ok := need[s[left]]; ok {
				if val == window[s[left]] {
					valid--
				}
				window[s[left]]--
			}

			left++
		}
	}

	if length == math.MaxInt16 {
		return ""
	}
	return s[start : start+length]
}
