package main

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	win := make(map[byte]int)
	valid, left, right := 0, 0, 0
	for i := range s1 {
		need[s1[i]]++
	}

	for right < len(s2) {
		if val, ok := need[s2[right]]; ok {
			win[s2[right]]++
			if val == win[s2[right]] {
				valid++
			}
		}

		right++

		for valid == len(need) {
			if left < right {
				if val, ok := need[s2[left]]; ok {
					if val == win[s2[left]] {
						if right-left == len(s1) {
							return true
						}
						valid--
					}
					win[s2[left]]--
				}

				left++
			}
		}
	}

	return false

}

func checkInclusion1(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	s1Len := len(s1)
	valid := 0
	left, right := 0, 0

	for i := range s1 {
		need[s1[i]]++
	}

	for right < len(s2) {
		if val, ok := need[s2[right]]; ok {
			window[s2[right]]++
			if val == window[s2[right]] {
				valid++
			}
		}
		right++

		for valid == len(need) {
			if right-left == s1Len {
				return true
			} else {
				remove := s2[left]
				if val, ok := need[remove]; ok {
					if val == window[remove] {
						valid--
					}
					window[remove]--
				}
				left++
			}
		}

	}

	return false
}
