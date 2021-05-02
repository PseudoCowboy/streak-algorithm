package main

func longestPalindrome(s string) int {
	m := make(map[byte]int)

	for i := range s {
		m[s[i]]++
	}

	sum := 0
	for _, v := range m {
		sum += (v / 2) * 2
	}

	if sum < len(s) {
		sum++
	}

	return sum
}
