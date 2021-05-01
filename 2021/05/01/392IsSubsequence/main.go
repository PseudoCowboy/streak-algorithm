package main

func isSubsequence(s string, t string) bool {
	for len(s) > 0 && len(t)-len(s) >= 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
