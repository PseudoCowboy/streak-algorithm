package main

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	window := make(map[byte]int)
	need := make(map[byte]int)
	left, right, valid := 0, 0, 0

	for i := range p {
		need[p[i]]++
	}

	for right < len(s) {
		if val, ok := need[s[right]]; ok {
			window[s[right]]++
			if val == window[s[right]] {
				valid++
			}
		}

		right++

		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
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
	return result
}
