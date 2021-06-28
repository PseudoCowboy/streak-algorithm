package main

func removeDuplicates(s string) string {
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(b) == 0 {
			b = append(b, s[i])
		} else {
			if b[len(b)-1] == s[i] {
				b = b[:len(b)-1]
			} else {
				b = append(b, s[i])
			}
		}
	}

	return string(b)
}
