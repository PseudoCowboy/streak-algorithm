package main

import "fmt"

func main() {
	a := "{}[]()"
	for i := range a {
		fmt.Println(a[i])
	}

	// fmt.Println(string(sb))
}

func reverseStr(s string, k int) string {
	b := make([]byte, 0)
	sb := []byte(s)
	count := 0
	for len(sb) != 0 {
		if count%2 == 0 {
			if len(sb) >= k {
				b = append(b, reverse(sb[:k])...)
				sb = sb[k:]
			} else {
				b = append(b, reverse(sb)...)
				sb = []byte{}
			}
		} else {
			if len(sb) >= k {
				b = append(b, sb[:k]...)
				sb = sb[k:]
			} else {
				b = append(b, sb...)
				sb = []byte{}
			}
		}
		count++
	}

	return string(b)
}

func reverse(s []byte) []byte {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

func reverseStr1(s string, k int) string {
	sb := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			reverse1(sb[i : i+k])
		} else {
			reverse1(sb[i:len(s)])
		}
	}

	return string(sb)
}

func reverse1(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
