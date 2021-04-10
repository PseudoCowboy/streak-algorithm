package main

func main() {
	s := "a"
	lengthOfLastWord(s)
}

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 && string(s[last]) == " " {
		last--
	}
	first := last
	for last >= 0 && string(s[last]) != " " {
		last--
	}
	return first - last
}
