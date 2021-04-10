package main

func main() {
	strStr("hello", "ll")
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}

	needleLength := len(needle)
	haystackLength := len(haystack)
	index := 0

	for index <= haystackLength-needleLength {
		for i := range needle {
			if needle[i] != haystack[i+index] {
				break
			}
			if i == needleLength-1 {
				return index
			}
		}
		index++
	}

	return -1
}
