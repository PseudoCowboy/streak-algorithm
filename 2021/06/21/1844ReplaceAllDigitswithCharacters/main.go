package main

func replaceDigits(s string) string {
	sb := make([]byte, 0)
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			sb = append(sb, s[i])
		} else {
			sb = append(sb, sb[i-1]+s[i]-'0')
		}
	}

	return string(sb)
}
