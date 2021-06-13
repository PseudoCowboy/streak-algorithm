package main

func makeEqual(words []string) bool {
	wl := len(words)
	sum := 0
	for i := range words {
		sum += len(words[i])
	}
	if sum%wl != 0 {
		return false
	}

	need := make(map[byte]int)

	for i := range words {
		for j := range words[i] {
			need[words[i][j]]++
		}
	}

	count := -1
	for _, v := range need {
		if count == -1 {
			count = v
		}
		if count%wl != 0 {
			return false
		}
	}

	return true

}
