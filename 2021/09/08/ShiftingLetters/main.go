package main

func shiftingLetters(s string, shifts []int) string {
	for i := len(shifts) - 2; i >= 0; i-- {
		shifts[i] = (shifts[i+1] + shifts[i]) % 26
	}
	sb := []byte(s)
	for i := range sb {
		sb[i] = 'a' + (sb[i]-'a'+byte(shifts[i]))%26
	}

	return string(sb)

}
