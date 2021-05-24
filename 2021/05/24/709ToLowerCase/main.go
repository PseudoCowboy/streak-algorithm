package main

func toLowerCase(s string) string {
	result := make([]rune, len(s))

	for i, v := range s {
		if v > 'A' && v < 'Z' {
			result[i] = v + rune(32)
		} else {
			result[i] = v
		}
	}

	return string(result)
}
