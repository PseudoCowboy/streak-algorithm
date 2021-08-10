package main

func minFlipsMonoIncr(s string) int {
	zero := 0
	for _, v := range s {
		if v == '0' {
			zero++
		}
	}
	curFlip := 0
	minFlip := len(s)
	for i := 0; i < len(s) && curFlip < minFlip; i++ {
		if s[i] == '0' {
			zero--
			continue
		}
		// Flip all '0' to '1'
		if curFlip+zero < minFlip {
			minFlip = curFlip + zero
		}
		// Flip S[i] to '0'
		curFlip++
	}
	if curFlip < minFlip {
		minFlip = curFlip
	}
	return minFlip
}
